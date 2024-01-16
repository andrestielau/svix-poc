package svix

import (
	"encoding/json"
	"log"
	"os"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/package/app/cmd"
	"svix-poc/package/app/flag"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	svix "github.com/svix/svix-webhooks/go"
)

var messageAppId string
var messageEventId string
var messagePayload string
var Message = cmd.New("message",
	cmd.Alias("m"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runListMessages), cmd.Flags(listFlagsOpt...)),
		cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateMessages), cmd.Flags(
			flag.StringP(&messageEventId, "event-type-id", "", "e", "event type id"),
			flag.StringP(&messagePayload, "payload", "", "p", "payload"),
		)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(runGetMessages)),
	),
	cmd.PFlags(
		flag.StringP(&messageAppId, "app", "", "a", "application"),
	),
)

func runListMessages(cmd *cobra.Command, _ []string) {
	res := lo.Must(svixclient.Client.Message.List(cmd.Context(), messageAppId, &svix.MessageListOptions{
		Iterator: lo.EmptyableToPtr(cursor),
		Limit:    &limit,
	}))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runCreateMessages(cmd *cobra.Command, _ []string) {
	var payload map[string]any
	lo.Must0(json.Unmarshal([]byte(messagePayload), &payload))
	m := &svix.MessageIn{
		EventType: messageEventId,
		Payload:   payload,
	}
	if messageEventId != "" {
		m.SetEventId(messageEventId)
	}
	res := lo.Must(svixclient.Client.Message.Create(cmd.Context(), messageAppId, m))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runGetMessages(cmd *cobra.Command, ids []string) {
	var res []svix.MessageOut
	var errs []error
	for _, id := range ids {
		if message, err := svixclient.Client.Message.Get(cmd.Context(), messageAppId, id); err != nil {
			errs = append(errs, err)
		} else {
			res = append(res, *message)
		}
	}
	log.Println("Errors: ", errs)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}
