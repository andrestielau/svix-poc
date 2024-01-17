package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/api/v1"
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
	for _, id := range ids {
		res := lo.Must(client.GetEndpointWithResponse(cmd.Context(), id, &webhooksv1.GetEndpointParams{}))
		log.Println(res)
	}
}
func runListEndpoints(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListEndpointsWithResponse(cmd.Context(), &webhooksv1.ListEndpointsParams{}))
	log.Println(res)
}
func runCreateEndpoints(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateEndpointsWithResponse(cmd.Context(), &webhooksv1.CreateEndpointsParams{}, []webhooksv1.NewEndpoint{}))
	log.Println(res)
}
