package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/lib/app/cmd"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Message = cmd.New("message", cmd.Add(
	cmd.New("get", cmd.Alias("g"), cmd.Run(runGetMessage)),
	cmd.New("list", cmd.Alias("l"), cmd.Run(runListMessages)),
	cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateMessages)),
))

func runGetMessage(cmd *cobra.Command, ids []string) {
	res := lo.Must(client.GetMessages(cmd.Context(), &webhooksv1.GetMessagesRequest{Ids: ids}))
	log.Println(res.Data)
}
func runListMessages(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListMessages(cmd.Context(), &webhooksv1.ListMessagesRequest{}))
	log.Println(res.Data)
}
func runCreateMessages(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateMessages(cmd.Context(), &webhooksv1.CreateMessagesRequest{}))
	log.Println(res.Data)
}
