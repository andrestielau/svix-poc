package tests_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(m *testing.M) { os.Exit(run(m)) }

var grpcClient webhooksv1.WebHookServiceClient

func run(m *testing.M) int {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
	}

	conn := lo.Must(grpc.Dial("localhost:4315", grpc.WithTransportCredentials(insecure.NewCredentials())))
	grpcClient = webhooksv1.NewWebHookServiceClient(conn)
	defer conn.Close()

	lo.Must(grpcClient.CreateApps(context.Background(), &webhooksv1.CreateAppsRequest{
		Data: []*webhooksv1.App{{
			Uid:  tenantId,
			Name: "Test App-" + tenantId,
		}},
	}))

	return m.Run()
}

func CheckReceived(t *testing.T, eventId string) {
	time.Sleep(15 * time.Second)
	res4, err := http.Get("http://localhost:3000/history?filter=/" + eventId)
	require.NoError(t, err)
	defer res4.Body.Close()
	var data []any
	err = json.NewDecoder(res4.Body).Decode(&data)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(data), 1)
	require.Equal(t, float64(200), data[0].(map[string]any)["response"].(map[string]any)["status"])
}

func CheckErr(t *testing.T, errs []*webhooksv1.Error) {
	if len(errs) > 0 {
		if errs[0].Reason != "409 Conflict" {
			require.Equal(t, []*webhooksv1.Error(nil), errs)
		}
	}
}

func Setup(t *testing.T, ctx context.Context, eventId string) {
	res, err := grpcClient.CreateEventTypes(ctx, &webhooksv1.CreateEventTypesRequest{
		Data: []*webhooksv1.EventType{{
			Name: "asd" + eventId,
		}},
	})
	require.NoError(t, err)
	CheckErr(t, res.Errors)
	endpointId := uuid.NewString()
	res2, err := grpcClient.CreateEndpoints(ctx, &webhooksv1.CreateEndpointsRequest{ // Register Endpoint in Svix
		TenantId: tenantId,
		Data: []*webhooksv1.Endpoint{{
			Uid:         endpointId,
			Url:         "http://smocker:8080/" + eventId,
			FilterTypes: []string{"asd" + eventId},
		}},
	})
	require.NoError(t, err)
	CheckErr(t, res2.Errors)
}
