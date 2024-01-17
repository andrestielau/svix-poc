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

var eventTypeName string
var EventType = cmd.New("eventType",
	cmd.Alias("et"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runListEventTypes), cmd.Flags(listFlagsOpt...)),
		cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEventTypes), cmd.Flags(
			flag.StringP(&eventTypeName, "name", "", "n", "name"),
		)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEventTypes)),
	),
)

func runListEventTypes(cmd *cobra.Command, _ []string) {
	res := lo.Must(svixclient.Client.EventType.List(cmd.Context(), &svix.EventTypeListOptions{
		Iterator: lo.EmptyableToPtr(cursor),
		Limit:    &limit,
	}))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runCreateEventTypes(cmd *cobra.Command, _ []string) {
	m := &svix.EventTypeIn{
		Name: eventTypeName,
	}
	res := lo.Must(svixclient.Client.EventType.Create(cmd.Context(), m))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runGetEventTypes(cmd *cobra.Command, ids []string) {
	var res []svix.EventTypeOut
	var errs []error
	for _, id := range ids {
		if eventType, err := svixclient.Client.EventType.Get(cmd.Context(), id); err != nil {
			errs = append(errs, err)
		} else {
			res = append(res, *eventType)
		}
	}
	log.Println("Errors: ", errs)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}
