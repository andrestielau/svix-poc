package tests

import "testing"

func TestTopic(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{
			name: "given that app and endpoint are properly configured, when receiving a message from pubsub, should send webhook",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// Create App
			// Create Endpoint
			// Start Mock
			// Send Message
			// Validate Reception
		})
	}
}
