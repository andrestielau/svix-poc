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

var endpointUid string
var endpointUrl string
var endpointSecret string
var endpointAppId string
var Endpoint = cmd.New("endpoint",
	cmd.Alias("e"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints), cmd.Flags(listFlagsOpt...)),
		cmd.New("create", cmd.Alias("c"), cmd.Run(runCreateEndpoints), cmd.Flags(
			flag.StringP(&endpointSecret, "secret", "", "s", "secret"),
			flag.StringP(&endpointUrl, "url", "", "u", "url"),
			flag.StringP(&endpointUid, "uid", "", "i", "uid"),
		)),
		cmd.New("get", cmd.Alias("g"), cmd.Run(runGetEndpoints)),
	),
	cmd.PFlags(
		flag.StringP(&endpointAppId, "app", "", "a", "application"),
	),
)

func runListEndpoints(cmd *cobra.Command, _ []string) {
	res := lo.Must(svixclient.Client.Endpoint.List(cmd.Context(), endpointAppId, &svix.EndpointListOptions{
		Iterator: lo.EmptyableToPtr(cursor),
		Limit:    &limit,
	}))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runCreateEndpoints(cmd *cobra.Command, _ []string) {
	m := &svix.EndpointIn{
		Url: endpointUrl,
	}
	if endpointSecret != "" {
		m.SetSecret(endpointSecret)
	}
	if endpointUid != "" {
		m.SetUid(endpointUid)
	}
	res := lo.Must(svixclient.Client.Endpoint.Create(cmd.Context(), endpointAppId, m))
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}

func runGetEndpoints(cmd *cobra.Command, ids []string) {
	var res []svix.EndpointOut
	var errs []error
	for _, id := range ids {
		if endpoint, err := svixclient.Client.Endpoint.Get(cmd.Context(), endpointAppId, id); err != nil {
			errs = append(errs, err)
		} else {
			res = append(res, *endpoint)
		}
	}
	log.Println("Errors: ", errs)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent(" ", " ")
	lo.Must0(e.Encode(res))
}
