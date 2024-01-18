package svix

import (
	"encoding/json"
	"log"
	"os"
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/lib/app/cmd"
	"svix-poc/lib/app/flag"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	svix "github.com/svix/svix-webhooks/go"
)

var attemptAppId string
var attemptMessageId string
var Attempt = cmd.New("message",
	cmd.Alias("r"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runListAttempts), cmd.Flags(listFlagsOpt...)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(runGetAttempts)),
	),
	cmd.PFlags(
		flag.StringP(&attemptAppId, "app", "", "a", "application id"),
		flag.StringP(&attemptMessageId, "message", "", "m", "message id"),
	),
)

func runListAttempts(cmd *cobra.Command, _ []string) {
	res := lo.Must(svixclient.Client.MessageAttempt.ListByMsg(cmd.Context(), attemptAppId, attemptMessageId, &svix.MessageAttemptListOptions{
		Iterator: lo.EmptyableToPtr(cursor),
		Limit:    &limit,
	}))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runGetAttempts(cmd *cobra.Command, ids []string) {
	var res []svix.MessageAttemptOut
	var errs []error
	for _, id := range ids {
		if message, err := svixclient.Client.MessageAttempt.Get(cmd.Context(), attemptAppId, attemptMessageId, id); err != nil {
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
