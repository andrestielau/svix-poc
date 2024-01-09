package call

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"svix-poc/cmd/call/app"
	eventtype "svix-poc/cmd/call/event-type"
	"svix-poc/internal/webhooks"
	"svix-poc/package/app/cfg"
	"svix-poc/package/app/cmd"
	"svix-poc/package/app/flag"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	svix "github.com/svix/svix-webhooks/go"
)

var Root = cmd.New("call",
	cmd.Alias("c"),
	cmd.PFlags(
		flag.StringP(&webhooks.DefaultAuthToken, "auth-token", "x", "t", "auth token"),
		flag.StringP(&webhooks.DefaultURL, "url", "http://localhost:8071", "u", "url"),
	),
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		webhooks.Init()
	}),
	cmd.Add(eventtype.Root, app.Root,
		cmd.New("demo",
			cmd.Alias("d"),
			cmd.Run(runDemo),
		),
	),
)

func runDemo(cmd *cobra.Command, _ []string) {
	var c webhooks.Config
	lo.Must0(cfg.ReadFile(".wh.yml", &c))

	wh := webhooks.Client
	ctx := cmd.Context()
	for _, event := range c.Events {
		res := must(wh.EventType.Create(ctx, &event))
		log.Println("Created Event Type", res)
	}
	log.Println(c.Apps)
	ids := map[string]string{}
	for _, app := range c.Apps {
		res := lo.Must(wh.Application.GetOrCreate(ctx, &app.Def))
		ids[strings.ToLower(app.Def.Name)] = res.Id
		log.Println("Created Application", res)
		for _, ep := range app.Endpoints {
			res2 := must(wh.Endpoint.Create(ctx, res.Id, &ep))
			log.Println("Created Endpoint", res2)
		}
	}

	var msgs map[string][]svix.MessageIn
	lo.Must0(cfg.ReadFile(".msg.yml", &msgs))
	log.Println(msgs)

	for app, m := range msgs {
		for _, msg := range m {
			log.Println(app, ids[app])
			res := must(wh.Message.Create(ctx, ids[app], &msg))
			log.Println("Created Message", res)
		}
	}
}
func must[T any](t T, err error) T {
	must0(err)
	return t
}
func must0(err error) {
	if err == nil {
		return
	}
	var serr *svix.Error
	if !errors.As(err, &serr) {
		panic(err)
	}
	if serr.Status() == http.StatusConflict {
		return
	}
	log.Println(string(serr.Body()))
	panic(serr.Error())
}
