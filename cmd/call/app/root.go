package app

import (
	"svix-poc/package/app/cmd"

	"svix-poc/cmd/call/app/attempt"
	"svix-poc/cmd/call/app/endpoint"

	"github.com/spf13/cobra"
)

var Root = cmd.New("app",
	cmd.Alias("a"),
	cmd.Add(
		cmd.New("create", cmd.Alias("c"), cmd.Run(CreateApplications)),
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListApplications)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetApplications)),
		cmd.New("update", cmd.Alias("u"), cmd.Run(UpdateApplications)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteApplications)),
		cmd.New("patch", cmd.Alias("p"), cmd.Run(PatchApplications)),
		endpoint.Root, attempt.Root,
	),
)

func CreateApplications(cmd *cobra.Command, _ []string) {

}

func ListApplications(cmd *cobra.Command, _ []string) {

}

func GetApplications(cmd *cobra.Command, _ []string) {

}

func UpdateApplications(cmd *cobra.Command, _ []string) {

}

func DeleteApplications(cmd *cobra.Command, _ []string) {

}

func PatchApplications(cmd *cobra.Command, _ []string) {

}
