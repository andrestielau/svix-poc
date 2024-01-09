package endpoint

import (
	"log"
	"svix-poc/package/app/cmd"

	"svix-poc/internal/webhooks"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	svix "github.com/svix/svix-webhooks/go"
)

var Root = cmd.New("endpoint",
	cmd.Alias("e"),
	cmd.Add(
		cmd.New("create", cmd.Alias("c"), cmd.Run(CreateEndpoints)),
		cmd.New("list", cmd.Alias("l"), cmd.Run(ListEndpoints)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(GetEndpoints)),
		cmd.New("update", cmd.Alias("u"), cmd.Run(UpdateEndpoints)),
		cmd.New("delete", cmd.Alias("d"), cmd.Run(DeleteEndpoints)),
		cmd.New("patch", cmd.Alias("p"), cmd.Run(PatchEndpoints)),
	),
)

func CreateEndpoints(cmd *cobra.Command, _ []string) {

}

func ListEndpoints(cmd *cobra.Command, apps []string) {
	for _, app := range apps {
		res := lo.Must(webhooks.Client.Endpoint.List(cmd.Context(), app, &svix.EndpointListOptions{}))
		log.Println(app, res)
	}
}

func GetEndpoints(cmd *cobra.Command, _ []string) {

}

func UpdateEndpoints(cmd *cobra.Command, _ []string) {

}

func DeleteEndpoints(cmd *cobra.Command, _ []string) {

}

func PatchEndpoints(cmd *cobra.Command, _ []string) {

}
