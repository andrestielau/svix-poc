package integration

import (
	"svix-poc/package/app/cmd"

	"github.com/spf13/cobra"
)

var Root = cmd.New("integration",
	cmd.Alias("i"),
	cmd.Add(
		cmd.New("create", cmd.Alias("c"), cmd.Run(CreateIntegrations)),
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListIntegrations)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetIntegrations)),
		cmd.New("update", cmd.Alias("u"), cmd.Run(UpdateIntegrations)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteIntegrations)),
		cmd.New("key", cmd.Alias("k"), cmd.Run(GetIntegrationKeys)),
		cmd.New("rotate", cmd.Alias("r"), cmd.Run(RotateIntegrationKeys)),
	),
)

func CreateIntegrations(cmd *cobra.Command, _ []string) {

}

func ListIntegrations(cmd *cobra.Command, _ []string) {

}

func GetIntegrations(cmd *cobra.Command, _ []string) {

}

func UpdateIntegrations(cmd *cobra.Command, _ []string) {

}

func DeleteIntegrations(cmd *cobra.Command, _ []string) {

}

func GetIntegrationKeys(cmd *cobra.Command, _ []string) {

}
func RotateIntegrationKeys(cmd *cobra.Command, _ []string) {

}
