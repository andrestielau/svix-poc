package routertopic

import (
	"svix-poc/app/router"
	"svix-poc/app/router/repo"
	"svix-poc/package/app"

	"github.com/google/wire"
)

func Provide(d router.Dependencies) *Handler {
	return &Handler{
		Dependencies: d,
		BaseActor: app.NewActor(app.Actors{
			repo.SingletonKey: d.Repository,
		}),
	}
}

var Set = wire.NewSet(Provide)
