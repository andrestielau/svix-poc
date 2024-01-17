package tests_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	webhooksgrpc "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestApi(t *testing.T) {
	for _, tt := range []struct {
		name string
		uid  string
	}{
		{
			name: "given that app and endpoint are properly configured, when sending a receiving from api, should send webhook",
			uid:  uuid.NewString(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tenantId := uuid.NewString()
			ctx := context.Background()
			client := lo.Must(webhooksv1.NewClient("http://localhost:2635"))
			grpcClient := webhooksgrpc.NewWebHookServiceClient(lo.Must(grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials()))))
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Log(utils.JsonReq[any](w, r))
			}))
			res, err := grpcClient.CreateApps(ctx, &webhooksgrpc.CreateAppsRequest{
				Data: []*webhooksgrpc.App{{
					Uid:  tenantId,
					Name: "Test App-" + tenantId,
				}},
			})
			require.NoError(t, err)
			t.Log(res)

			res2, err := client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsParams{ // Create Endpoint
				TenantId: tt.uid,
			}, []webhooksv1.NewEndpoint{{
				Url: server.URL,
			}})
			require.NoError(t, err)
			t.Log(res2)
			res3, err := client.CreateMessages(ctx, &webhooksv1.CreateMessagesParams{ // Send Message
				TenantId: tt.uid,
			}, []webhooksv1.NewMessage{{
				Payload: `{ "foo": "bar" }`,
			}})
			require.NoError(t, err)
			t.Log(res3)
			// Wait a bit
			time.Sleep(5 * time.Second)
		})
	}
}
