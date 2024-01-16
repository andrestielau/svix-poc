package webhook

import (
	svixclient "svix-poc/app/webhook/svix"

	"github.com/google/wire"
)

type Dependencies struct {
	*svixclient.SvixClient
}

var Set = wire.NewSet(
	svixclient.Provide,
	wire.Struct(new(Dependencies), "*"),
)
