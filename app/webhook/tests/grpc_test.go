package tests_test

import (
	"context"
	"encoding/json"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
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
			eventId := uuid.NewString()
			Setup(t, ctx, eventId)
			res4, err := grpcClient.CreateMessages(ctx, &webhooksv1.CreateMessagesRequest{
				TenantId: tenantId,
				Data: []*webhooksv1.Message{{
					EventType: "asd" + eventId,
					EventId:   eventId,
					Payload: lo.Must(json.Marshal(map[string]any{
						"foo": "grpc",
					})),
				}},
			})
			require.NoError(t, err)
			CheckErr(t, res4.Errors)
			CheckReceived(t, eventId)
		})
	}
}
