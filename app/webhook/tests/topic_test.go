package tests_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/utils"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestTopic(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{
			name: "given that app and endpoint are properly configured, when receiving a message from pubsub, should send webhook",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tenantId := uuid.NewString()
			ctx := context.Background()
			client := webhooksv1.NewWebHookServiceClient(lo.Must(grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials()))))
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Log(utils.JsonReq[any](w, r))
			}))
			res, err := client.CreateApps(ctx, &webhooksv1.CreateAppsRequest{
				Data: []*webhooksv1.App{{
					Uid:  tenantId,
					Name: "Test App-" + tenantId,
				}},
			})
			require.NoError(t, err)
			t.Log(res)
			res2, err := client.CreateEventTypes(ctx, &webhooksv1.CreateEventTypesRequest{
				Data: []*webhooksv1.EventType{{
					Name: "asd",
				}},
			})
			require.NoError(t, err)
			t.Log(res2)
			endpointId := uuid.NewString()
			res3, err := client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsRequest{ // Register Endpoint in Svix
				TenantId: tenantId,
				Data: []*webhooksv1.Endpoint{{
					Uid: endpointId,
					Url: server.URL,
				}},
			})
			require.NoError(t, err)
			t.Log(res3)
			publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
				ProjectID: "demo",
			}, watermill.NewStdLogger(true, false))
			require.NoError(t, err)
			// Publish Message
			err = publisher.Publish("webhook", &message.Message{
				Metadata: map[string]string{
					"app":       tenantId,
					"EventType": "asd",
				},
				Payload: []byte(`{ "foo": "bar" }`),
			})
			require.NoError(t, err)
			// Wait a bit
			time.Sleep(5 * time.Second)
		})
	}
}
