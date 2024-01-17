package tests_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpc(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{
			name: "given that app and endpoint are properly configured, when receiving a message from grpc, should send webhook",
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
			endpointId := uuid.NewString()
			res2, err := client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsRequest{ // Register Endpoint in Svix
				TenantId: tenantId,
				Data: []*webhooksv1.Endpoint{{
					Uid: endpointId,
					Url: server.URL,
				}},
			})
			require.NoError(t, err)
			t.Log(res2)
			res3, err := client.CreateMessages(ctx, &webhooksv1.CreateMessagesRequest{
				TenantId: tenantId,
				Data: []*webhooksv1.Message{{
					EventType: "foo",
					Id:        "asd",
					EventId:   "asdadsdas",
					Payload: lo.Must(json.Marshal(map[string]any{
						"foo": "bar",
					})),
				}},
			})
			require.NoError(t, err)
			t.Log(res3)
			// Wait a bit
			time.Sleep(5 * time.Minute)
		})
	}
}
