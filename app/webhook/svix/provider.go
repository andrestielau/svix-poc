package svixclient

import (
	"context"
	"svix-poc/package/app"

	svix "github.com/svix/svix-webhooks/go"
)

type Provider struct {
	*app.BaseActor
	*svix.Svix
}

func Provide() *Provider {
	return &Provider{
		BaseActor: app.NewActor(),
	}
}
func (p *Provider) Start(ctx context.Context) (first bool, err error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if first, err = p.BaseStart(ctx); first || err != nil {
		return first, err
	} else if p.Svix, err = New(
		ProvideURL(),
		ProvideAuthToken(),
	); err != nil {
		return true, err
	}
	return true, nil
}
