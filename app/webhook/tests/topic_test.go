package tests_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopic(t *testing.T) {
	for _, tt := range []struct {
		name string
	}{
		{
			name: "given that app and endpoint are properly configured, when receiving a message from pubsub, should send webhook",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}))
			log.Println(server)
			// Create App
			// Create Endpoint
			// Start Mock
			// Send Message
			// Validate Reception
		})
	}
}
