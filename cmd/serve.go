package cmd

import (
	"log"
	"svix-poc/x/app/cmd"

	"github.com/spf13/cobra"
)

var Serve = cmd.New("serve",
	cmd.Alias("s"),
	cmd.Run(runServe),
)

func runServe(cmd *cobra.Command, modules []string) {
	if len(modules) == 0 {
		modules = []string{"email", "webhook", "router"}
	}
	log.Println(modules)
	for _, module := range modules {
		switch module {
		case "email":
		case "webhook":
		case "router":
		default:
		}
	}
	
}
