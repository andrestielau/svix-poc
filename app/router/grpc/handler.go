package routergrpc

import (
	"context"
	"log"
	"net"
	"strings"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/grpc/v1"
	"svix-poc/package/app"
	schemav1 "svix-poc/package/schema/grpc/v1"

	"google.golang.org/grpc"
)

type Handler struct {
	router.Dependencies
	*app.BaseActor
	Server *grpc.Server
}

var _ eventsv1.EventServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}

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
		log.Println("router grpc listening on http://" + host)
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
