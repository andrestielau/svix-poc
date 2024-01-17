package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/api/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var EventType = cmd.New("type", cmd.Add(
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListEventTypes)),
))

func runListEventTypes(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListTypesWithResponse(cmd.Context(), &webhooksv1.ListTypesParams{}))
	log.Println(res)
}
