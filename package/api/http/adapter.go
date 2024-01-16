package http

import (
	"context"
	"log"
	"net/http"
	"svix-poc/package/app"
)

type AdapterOptions struct {
	Addr    string
	Handler http.Handler
}
type Adapter struct {
	AdapterOptions
	*app.BaseActor
	s *http.Server
}

func NewAdapter(opts AdapterOptions, children app.Actors) *Adapter {
	return &Adapter{
		AdapterOptions: opts,
		BaseActor:      app.NewActor(children),
	}
}
func (h *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	h.s = &http.Server{
		Addr:    h.Addr,
		Handler: h.Handler,
	}
	go func() {
		if err := h.s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	return true, nil
}
func (h *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	return true, h.s.Close()
}
