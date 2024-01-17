package tests_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	svixclient "svix-poc/app/webhook/svix"
	"testing"

	"github.com/google/uuid"
	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

func TestApi(t *testing.T) {
	for _, tt := range []struct {
		name string
		uid  string
	}{
		{
			name: "given that app and endpoint are properly configured, when sending a receiving from api, should send webhook",
			uid:  uuid.NewString(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := lo.Must(webhooksv1.NewClient(""))
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}))
			appReq := &svix.ApplicationIn{Name: "App_" + tt.uid}
			appReq.SetUid(tt.uid)
			lo.Must(svixclient.Client.Application.Create(ctx, appReq))             // Create App
			lo.Must(client.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsParams{ // Create Endpoint
				TenantId: tt.uid,
			}, []webhooksv1.NewEndpoint{{
				Url: server.URL,
			}}))
			lo.Must(client.CreateMessages(ctx, &webhooksv1.CreateMessagesParams{ // Send Message
				TenantId: tt.uid,
			}, []webhooksv1.NewMessage{{
				Payload: `{ "foo": "bar" }`,
			}}))
			// Wait a bit
		})
	}
}
