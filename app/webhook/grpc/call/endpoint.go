package call

import (
	"svix-poc/lib/app/cmd"

	"github.com/spf13/cobra"
)

var Endpoint = cmd.New("endpoint", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEndpoint)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEndpoints)),
))

func runGetEndpoint(_ *cobra.Command, _ []string) {

}
func runListEndpoints(_ *cobra.Command, _ []string) {

}
func runCreateEndpoints(_ *cobra.Command, _ []string) {

}
