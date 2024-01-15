package cmd

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"svix-poc/x/app/cfg"
	"svix-poc/x/app/cmd"
	"sync"
	"syscall"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	svix "github.com/svix/svix-webhooks/go"
)

var Mock = cmd.New("mock",
	cmd.Alias("m"),
	cmd.Run(runMock),
)

func runMock(cmd *cobra.Command, _ []string) {
	var c Config
	lo.Must0(cfg.ReadFile(".wh.yml", &c))

	wg := &sync.WaitGroup{}
	var ls []net.Listener
	defer func() {
		for i := range ls {
			go func(i int) { ls[i].Close() }(i)
		}
		wg.Wait()
	}()
	for _, app := range c.Apps {
		for _, ep := range app.Endpoints {
			l := lo.Must(net.Listen("tcp", strings.TrimPrefix(ep.Url, "http://127.0.0.1")))
			ls = append(ls, l)
			wg.Add(1)
			go func(l net.Listener, ep svix.EndpointIn) {
				defer wg.Done()
				func() {
					if err := recover(); err != nil {
						log.Println("Server crash", err)
					}
				}()
				var secret string
				if ep.Secret.IsSet() {
					secret = *ep.Secret.Get()
				}
				wh, err := svix.NewWebhook(secret)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("running on " + ep.Url)
				if err := http.Serve(l, &Handler{Events: ep.FilterTypes, WH: wh}); err != nil {
					log.Println("Err: ", err)
				}
			}(l, ep)
		}
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

type Handler struct {
	Events []string
	WH     *svix.Webhook
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Request received", req.Header)
	if req.Body != nil {
		defer req.Body.Close()
		payload := lo.Must(io.ReadAll(req.Body))
		log.Println("Body: ", string(payload))
		if err := h.WH.Verify(payload, req.Header); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	res.WriteHeader(http.StatusOK)
}
