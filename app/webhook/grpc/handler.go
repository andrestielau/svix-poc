package webhookgrpc

import (
	"context"
	"log"
	"net"
	"strings"
	"svix-poc/app/webhook"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"
	"svix-poc/package/app"
	schemav1 "svix-poc/package/schema/grpc/v1"

	"google.golang.org/grpc"
)

type Handler struct {
	webhook.Dependencies
	*app.BaseActor
	Server *grpc.Server
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
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

// CreateApps implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateApps(context.Context, *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	panic("unimplemented")
}
