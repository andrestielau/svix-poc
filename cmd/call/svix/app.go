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

var appUid string
var appName string
var App = cmd.New("app",
	cmd.Alias("a"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runListApps), cmd.Flags(listFlagsOpt...)),
		cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateApps), cmd.Flags(
			flag.StringP(&appName, "name", "", "n", "name"),
			flag.StringP(&appUid, "uid", "", "i", "uid"),
		)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(runGetApps)),
	),
)

func runListApps(cmd *cobra.Command, _ []string) {
	res := lo.Must(svixclient.Client.Application.List(cmd.Context(), &svix.ApplicationListOptions{
		Iterator: lo.EmptyableToPtr(cursor),
		Limit:    &limit,
	}))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runCreateApps(cmd *cobra.Command, _ []string) {
	m := &svix.ApplicationIn{
		Name: appName,
	}
	if appUid != "" {
		m.SetUid(appUid)
	}
	res := lo.Must(svixclient.Client.Application.Create(cmd.Context(), m))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runGetApps(cmd *cobra.Command, ids []string) {
	var res []svix.ApplicationOut
	var errs []error
	for _, id := range ids {
		if app, err := svixclient.Client.Application.Get(cmd.Context(), id); err != nil {
			errs = append(errs, err)
		} else {
			res = append(res, *app)
		}
	}
	log.Println("Errors: ", errs)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}
