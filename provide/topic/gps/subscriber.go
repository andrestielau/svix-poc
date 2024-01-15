package gps

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
)

func NewSubscriber(logger watermill.LoggerAdapter) (*googlecloud.Subscriber, error) {
	return googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return "test-sub_" + topic
			},
			ProjectID: "test-project",
		},
		logger,
	)
}
