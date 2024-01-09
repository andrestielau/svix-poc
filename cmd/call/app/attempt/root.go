package attempt

import (
	"encoding/json"
	"os"
	"svix-poc/package/app/cmd"

	"svix-poc/internal/webhooks"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Root = cmd.New("attempt",
	cmd.Alias("a"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListAttempt)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetAttempt)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteAttempt)),
		cmd.New("resend", cmd.Alias("r"), cmd.Run(ResendAttempt)),
	),
)

func ListAttempt(cmd *cobra.Command, apps []string) {
	ctx := cmd.Context()
	for _, app := range apps {
		eps := lo.Must(webhooks.Client.Endpoint.List(ctx, app, nil))
		for _, ep := range eps.Data {
			res := lo.Must(webhooks.Client.MessageAttempt.ListByEndpoint(ctx, app, ep.Id, nil))
			e := json.NewEncoder(os.Stdout)
			e.SetIndent(" ", " ")
			e.Encode(res)
		}
	}
}

func GetAttempt(cmd *cobra.Command, _ []string) {

}

func DeleteAttempt(cmd *cobra.Command, _ []string) {

}

func ResendAttempt(cmd *cobra.Command, _ []string) {

}
