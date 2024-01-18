package tests_test

import (
	"context"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	"testing"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
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
			eventId := uuid.NewString()
			Setup(t, ctx, eventId)
			_, err := client.CreateMessages(ctx, &webhooksv1.CreateMessagesParams{ // Send Message
				TenantId: tenantId,
			}, []webhooksv1.NewMessage{{
				EventType: "asd" + eventId,
				EventId:   lo.ToPtr(eventId),
				Payload:   `{ "foo": "topic" }`,
			}})
			require.NoError(t, err)
			CheckReceived(t, eventId)
		})
	}
}
