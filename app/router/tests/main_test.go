package tests

import (
	"os"
	eventsv1 "svix-poc/app/router/grpc/v1"
	"testing"

	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(m *testing.M) { os.Exit(run(m)) }

var grpcClient eventsv1.EventServiceClient

func run(m *testing.M) int {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
	}

	conn := lo.Must(grpc.Dial("localhost:2573", grpc.WithTransportCredentials(insecure.NewCredentials())))
	grpcClient = eventsv1.NewEventServiceClient(conn)
	defer conn.Close()

	return m.Run()
}
