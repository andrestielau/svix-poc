package tests_test

import (
	"context"
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"

	"github.com/samber/lo"
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
			ctx := context.Background()
			client := webhooksv1.NewWebHookServiceClient(lo.Must(grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials()))))
			/*server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			})) */
			res := lo.Must(client.CreateApps(ctx, &webhooksv1.CreateAppsRequest{}))
			log.Println(res)
			// Create Endpoint
			// Send Message
			// Wait a bit
		})
	}
}
