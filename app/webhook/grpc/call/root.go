package call

import (
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var client webhooksv1.WebHookServiceClient
var Root = cmd.New("webhook",
	cmd.Add(EventType, App, Endpoint, Message, Attempt),
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		client = webhooksv1.NewWebHookServiceClient(lo.Must(grpc.Dial("localhost:")))
	}),
)
