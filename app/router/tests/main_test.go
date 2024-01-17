package tests

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/docker/docker/client"
	"github.com/samber/lo"
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
