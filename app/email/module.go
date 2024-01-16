package email

import (
	"net/smtp"

	"github.com/google/wire"
)

type Dependencies struct {
	Client smtp.Client
}

var Set = wire.NewSet(
	wire.Struct(new(Dependencies), "*"),
)
