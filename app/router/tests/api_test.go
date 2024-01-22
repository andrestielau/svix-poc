package tests

import (
	"context"
	eventsv1 "svix-poc/app/router/api/v1"
	eventsgrpc "svix-poc/app/router/grpc/v1"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
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
			client := lo.Must(eventsv1.NewClient("http://localhost:3524"))
			eventId := uuid.NewString()
			_, err := grpcClient.CreateEventTypes(ctx, &eventsgrpc.CreateEventTypesRequest{
				Data: []*eventsgrpc.EventType{{
					Id: "dsa" + eventId,
				}},
			})
			require.NoError(t, err)
			tenantId := uuid.NewString()
			_, err = client.CreateSubscriptions(ctx, &eventsv1.CreateSubscriptionsParams{
				TenantId: tenantId,
			}, []eventsv1.NewSubscription{{
				EventTypeId: "dsa" + eventId,
			}})
			require.NoError(t, err)
			// Subscribe to Topic
			publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
				ProjectID: "demo",
			}, watermill.NewStdLogger(true, false))
			require.NoError(t, err)
			err = publisher.Publish("example1", &message.Message{ // Publish Message
				Metadata: map[string]string{
					"app":       tenantId,
					"EventType": "asd" + eventId,
				},
				Payload: []byte(`{ "foo": "http" }`),
			})
			require.NoError(t, err)
			time.Sleep(5 * time.Second) // Wait a bit
		})
	}
}
