package message

import (
	"svix-poc/package/app/cmd"

	"github.com/spf13/cobra"
)

var Root = cmd.New("message",
	cmd.Alias("m"),
	cmd.Add(
		cmd.New("create", cmd.Alias("c"), cmd.Run(CreateMessages)),
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListMessages)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetMessages)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteMessages)),
	),
)

func CreateMessages(cmd *cobra.Command, _ []string) {

}

func ListMessages(cmd *cobra.Command, _ []string) {

}

func GetMessages(cmd *cobra.Command, _ []string) {

}

func DeleteMessages(cmd *cobra.Command, _ []string) {

}
