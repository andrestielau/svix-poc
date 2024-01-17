package call

import (
	"log"
	webhooksv1 "svix-poc/app/webhook/api/v1"
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
	for _, id := range ids {
		res := lo.Must(client.GetMessageWithResponse(cmd.Context(), id, &webhooksv1.GetMessageParams{}))
		log.Println(res)
	}
}
func runListMessages(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.ListMessagesWithResponse(cmd.Context(), &webhooksv1.ListMessagesParams{}))
	log.Println(res)
}
func runCreateMessages(cmd *cobra.Command, _ []string) {
	res := lo.Must(client.CreateMessagesWithResponse(cmd.Context(), &webhooksv1.CreateMessagesParams{}, []webhooksv1.NewMessage{}))
	log.Println(res)
}
