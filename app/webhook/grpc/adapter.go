package webhookgrpc

import (
	"context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

// TODO: make generic for grpc

func (h *Handler) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	h.Server = grpc.NewServer()
	host := string(ProvideHost())
	l, err := net.Listen("tcp", host)
	if err != nil {
		return true, err
	}
	go func() {
		if strings.HasPrefix(host, ":") {
			host = "localhost" + host
		}
		log.Println("webhook grpc listening on http://" + host)
		if err := h.Server.Serve(l); err != nil {
			log.Println(err)
		}
	}()
	return true, nil
}
func (h *Handler) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	if h.Server == nil {
		return true, nil
	}
	h.Server.GracefulStop()
	return true, nil
}
