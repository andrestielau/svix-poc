package eventtype

import (
	"svix-poc/package/app/cmd"

	"github.com/spf13/cobra"
)

var Root = cmd.New("event-type",
	cmd.Alias("e"),
	cmd.Add(
		cmd.New("create", cmd.Alias("c"), cmd.Run(CreateEventTypes)),
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListEventTypes)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetEventTypes)),
		cmd.New("update", cmd.Alias("u"), cmd.Run(UpdateEventTypes)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteEventTypes)),
		cmd.New("patch", cmd.Alias("p"), cmd.Run(PatchEventTypes)),
	),
)

func CreateEventTypes(cmd *cobra.Command, _ []string) {

}

func ListEventTypes(cmd *cobra.Command, _ []string) {

}

func GetEventTypes(cmd *cobra.Command, _ []string) {

}

func UpdateEventTypes(cmd *cobra.Command, _ []string) {

}

func DeleteEventTypes(cmd *cobra.Command, _ []string) {

}

func PatchEventTypes(cmd *cobra.Command, _ []string) {

}
