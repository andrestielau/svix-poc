package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var App = cmd.New("app", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetApp)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListApps)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateApps)),
))

func runGetApp(cmd *cobra.Command, ids []string) {
	res := lo.Must(client.GetApps(cmd.Context(), &webhooksv1.GetAppsRequest{Ids: ids}))
	log.Println(res.Data)
}
func runListApps(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListApps(cmd.Context(), &webhooksv1.ListAppsRequest{}))
	log.Println(res.Data)
}
func runCreateApps(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateApps(cmd.Context(), &webhooksv1.CreateAppsRequest{}))
	log.Println(res.Data)
}
