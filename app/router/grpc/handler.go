package routergrpc

import (
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/grpc/v1"
	schemav1 "svix-poc/package/schema/grpc/v1"
)

type Handler struct {
	router.Dependencies
}

var _ eventsv1.EventServiceServer = &Handler{}
var _ schemav1.SchemaServiceServer = &Handler{}
