package gps

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
)

func NewPublisher(logger watermill.LoggerAdapter) (*googlecloud.Publisher, error) {
	return googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: "test-project",
	}, logger)
}
