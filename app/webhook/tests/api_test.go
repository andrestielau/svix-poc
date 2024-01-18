package tests_test

import (
	"context"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	webhooksgrpc "svix-poc/app/webhook/grpc/v1"
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
	}{
		{
			name: "given that app and endpoint are properly configured, when sending a receiving from api, should send webhook",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := lo.Must(webhooksv1.NewClient("http://localhost:2635"))
			conn, err := grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			grpcClient := webhooksgrpc.NewWebHookServiceClient(conn)

			_, err = grpcClient.CreateApps(ctx, &webhooksgrpc.CreateAppsRequest{
				Data: []*webhooksgrpc.App{{
					Uid:  tenantId,
					Name: "Test App-" + tenantId,
				}},
			})
			require.NoError(t, err)
			eventId := uuid.NewString()
			endpointId := uuid.NewString()
			_, err = client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsParams{ // Create Endpoint
				TenantId: tenantId,
			}, []webhooksv1.NewEndpoint{{
				Uid: lo.ToPtr(endpointId),
				Url: "http://smocker:8080/" + eventId,
			}})
			require.NoError(t, err)
			_, err = client.CreateMessages(ctx, &webhooksv1.CreateMessagesParams{ // Send Message
				TenantId: tenantId,
			}, []webhooksv1.NewMessage{{
				EventType: "asd",
				EventId:   lo.ToPtr(eventId),
				Payload:   `{ "foo": "` + time.Now().String() + `" }`,
			}})
			require.NoError(t, err)
			CheckReceived(t, eventId)
		})
	}
}
