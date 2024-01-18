package tests_test

import (
	"context"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const tenantId = "test-tenant"

func TestTopic(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{
			name: "given that app and endpoint are properly configured, when receiving a message from pubsub, should send webhook",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			conn, err := grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			client := webhooksv1.NewWebHookServiceClient(conn)
			res, err := client.CreateApps(ctx, &webhooksv1.CreateAppsRequest{
				Data: []*webhooksv1.App{{
					Uid:  tenantId,
					Name: "Test App-" + tenantId,
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res.Errors)
			eventId := uuid.NewString()
			res2, err := client.CreateEventTypes(ctx, &webhooksv1.CreateEventTypesRequest{
				Data: []*webhooksv1.EventType{{
					Name: "asd" + eventId,
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res2.Errors)
			endpointId := uuid.NewString()
			res3, err := client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsRequest{ // Register Endpoint in Svix
				TenantId: tenantId,
				Data: []*webhooksv1.Endpoint{{
					Uid: endpointId,
					Url: "http://smocker:8080/" + eventId,
					FilterTypes: []string{"asd" + eventId},
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res3.Errors)
			publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
				ProjectID: "demo",
			}, watermill.NewStdLogger(true, false))
			require.NoError(t, err)
			// Publish Message
			err = publisher.Publish("webhook", &message.Message{
				Metadata: map[string]string{
					"app":       tenantId,
					"EventType": "asd" + eventId,
				},
				Payload: []byte(`{ "foo": "bar" }`),
			})
			require.NoError(t, err)
			CheckReceived(t, eventId)
		})
	}
}
