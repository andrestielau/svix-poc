package svix

import (
	svixclient "svix-poc/app/webhook/svix"
	"svix-poc/lib/app/cmd"
	"svix-poc/lib/app/flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var limit int32
var cursor string
var listFlagsOpt = []func(*pflag.FlagSet){
	flag.Int32P(&limit, "limit", 10, "l", "limit"),
	flag.StringP(&cursor, "cursor", "", "c", "cursor"),
}
var Root = cmd.New("svix",
	cmd.Add(App, Endpoint, EventType, Message, Attempt), cmd.PFlags(
		flag.StringP(&svixclient.DefaultAuthToken, "auth-token", "x", "t", "auth token"),
		flag.StringP(&svixclient.DefaultURL, "svix-url", "http://localhost:8071", "s", "svix-url"),
	), cmd.PPreRun(func(cmd *cobra.Command, _ []string) { svixclient.Init() }),
)
