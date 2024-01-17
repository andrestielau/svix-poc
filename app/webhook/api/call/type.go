package call

import (
	"svix-poc/lib/app/cmd"

	"github.com/spf13/cobra"
)

var EventType = cmd.New("type", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEventType)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListEventTypes)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEventTypes)),
))

func runGetEventType(_ *cobra.Command, _ []string) {

}
func runListEventTypes(_ *cobra.Command, _ []string) {

}
func runCreateEventTypes(_ *cobra.Command, _ []string) {

}
