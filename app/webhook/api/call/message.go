package call

import (
	"svix-poc/lib/app/cmd"

	"github.com/spf13/cobra"
)

var Message = cmd.New("message", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetMessage)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListMessages)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateMessages)),
))

func runGetMessage(_ *cobra.Command, _ []string) {

}
func runListMessages(_ *cobra.Command, _ []string) {

}
func runCreateMessages(_ *cobra.Command, _ []string) {

}
