package tests_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"testing"
	"time"

	"github.com/docker/docker/client"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func TestMain(m *testing.M) { os.Exit(run(m)) }

func run(m *testing.M) int {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
	}

	if isIntegration() { // TODO: make these compatible with local testing
		cli := lo.Must(client.NewClientWithOpts())
		defer cli.Close()
		ctx := context.Background()
		compose := lo.Must(tc.NewDockerCompose("../../../compose.yml"))
		defer compose.Down(ctx, tc.RemoveOrphans(true), tc.RemoveVolumes(true), tc.RemoveImagesLocal)
		lo.Must0(compose.Up(ctx, tc.RunServices("postgres", "pgbouncer", "redis", "svix", "pubsub", "webhooks")))
	}
	return m.Run()
}

func isIntegration() bool {
	ev := strings.ToLower(os.Getenv("INTEGRATION_TEST"))
	return ev == "true" || ev == "t" || ev == "1"
}

func CheckReceived(t *testing.T, eventId string) {
	time.Sleep(10 * time.Second)
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
