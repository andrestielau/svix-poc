package server

import (
	"context"
	"log"
	"net"
	"svix-poc/lib/app"

	"google.golang.org/grpc"
)

type AdapterOptions struct {
	Addr     string
	Register func(*grpc.Server)
}
type Adapter struct {
	AdapterOptions
	*app.BaseActor
	s *grpc.Server
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
	h.s = grpc.NewServer()
	h.Register(h.s)
	l, err := net.Listen("tcp", h.Addr)
	if err != nil {
		return true, err
	}
	go func() {
		if err := h.s.Serve(l); err != nil {
			log.Println("Grpc Adapter", err)
		}
	}()
	return true, nil
}
func (h *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	h.s.GracefulStop()
	return true, nil
}
