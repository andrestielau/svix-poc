package email

import (
	"net/smtp"
	"svix-poc/package/app"

	"github.com/google/wire"
)

type Dependencies struct {
	Client smtp.Client
}
type Module struct {
	*app.BaseActor
	Dependencies
}

func New(d Dependencies) *Module {
	return &Module{
		BaseActor:    app.NewActor(),
		Dependencies: d,
	}
}

var Set = wire.NewSet(
	wire.Struct(new(Dependencies), "*"),
)
