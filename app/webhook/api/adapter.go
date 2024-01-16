package webhookapi

import (
	"context"
	"log"
	"net/http"
	"strings"
	webhookv1 "svix-poc/app/webhook/api/v1"
)

// TODO: make generic for http

func (h *Handler) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	host := string(ProvideHost())
	h.server = &http.Server{
		Addr:    host,
		Handler: webhookv1.Handler(h),
	}
	go func() {
		if strings.HasPrefix(host, ":") {
			host = "localhost" + host
		}
		log.Println("webhook api listening on http://" + host)
		if err := h.server.ListenAndServe(); err != nil {
			log.Println("err", err)
		}
	}()
	return true, nil
}
func (h *Handler) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	return true, h.server.Close()
}
