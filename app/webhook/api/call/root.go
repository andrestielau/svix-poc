package call

import (
	webhooksv1 "svix-poc/app/webhook/api/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var client webhooksv1.ClientWithResponsesInterface
var Root = cmd.New("webhook",
	cmd.Add(EventType, Endpoint, Message, Attempt),
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		client = lo.Must(webhooksv1.NewClientWithResponses("localhost:"))
	}),
)
