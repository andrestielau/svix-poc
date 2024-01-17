package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Endpoint = cmd.New("endpoint", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEndpoint)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEndpoints)),
))

func runGetEndpoint(cmd *cobra.Command, ids []string) {
	res := lo.Must(client.GetEndpoints(cmd.Context(), &webhooksv1.GetEndpointsRequest{Ids: ids}))
	log.Println(res.Data)
}
func runListEndpoints(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListEndpoints(cmd.Context(), &webhooksv1.ListEndpointsRequest{}))
	log.Println(res.Data)
}
func runCreateEndpoints(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateEndpoints(cmd.Context(), &webhooksv1.CreateEndpointsRequest{}))
	log.Println(res.Data)
}
