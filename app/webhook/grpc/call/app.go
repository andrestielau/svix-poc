package call

import (
	"svix-poc/lib/app/cmd"

	"github.com/spf13/cobra"
)

var App = cmd.New("app", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetApp)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListApps)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateApps)),
))

func runGetApp(_ *cobra.Command, _ []string) {

}
func runListApps(_ *cobra.Command, _ []string) {

}
func runCreateApps(_ *cobra.Command, _ []string) {

}
