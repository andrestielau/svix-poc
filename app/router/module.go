package router

import (
	"svix-poc/app/router/repo"

	"github.com/google/wire"
)

type Dependencies struct {
	*repo.Repository
}

var Set = wire.NewSet(
	repo.Provide,
	wire.Struct(new(Dependencies), "*"),
)
