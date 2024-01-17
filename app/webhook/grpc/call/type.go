package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var EventType = cmd.New("type", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEventType)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListEventTypes)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEventTypes)),
))

func runGetEventType(cmd *cobra.Command, ids []string) {
	res := lo.Must(client.GetEventTypes(cmd.Context(), &webhooksv1.GetEventTypesRequest{Ids: ids}))
	log.Println(res.Data)
}
func runListEventTypes(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListEventTypes(cmd.Context(), &webhooksv1.ListEventTypesRequest{}))
	log.Println(res.Data)
}
func runCreateEventTypes(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateEventTypes(cmd.Context(), &webhooksv1.CreateEventTypesRequest{}))
	log.Println(res.Data)
}
