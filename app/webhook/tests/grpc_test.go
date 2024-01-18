package tests_test

import (
	"context"
	"encoding/json"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"

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
			res2, err := client.CreateEventTypes(ctx, &webhooksv1.CreateEventTypesRequest{
				Data: []*webhooksv1.EventType{{
					Name: "asd",
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res2.Errors)
			eventId := uuid.NewString()
			endpointId := uuid.NewString()
			res3, err := client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsRequest{ // Register Endpoint in Svix
				TenantId: tenantId,
				Data: []*webhooksv1.Endpoint{{
					Uid: endpointId,
					Url: "http://smocker:8080/" + eventId,
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res3.Errors)
			res4, err := client.CreateMessages(ctx, &webhooksv1.CreateMessagesRequest{
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
			CheckErr(t, res4.Errors)
			CheckReceived(t, eventId)
		})
	}
}
