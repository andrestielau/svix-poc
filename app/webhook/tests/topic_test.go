package tests_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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
			eventId := uuid.NewString()
			Setup(t, ctx, eventId)
			publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
				ProjectID: "demo",
			}, watermill.NewStdLogger(true, false))
			require.NoError(t, err)

			err = publisher.Publish("webhook", &message.Message{ // Publish Message
				Metadata: map[string]string{
					"app":       tenantId,
					"EventType": "asd" + eventId,
				},
				Payload: []byte(`{ "foo": "http" }`),
			})
			require.NoError(t, err)
			CheckReceived(t, eventId)
		})
	}
}
