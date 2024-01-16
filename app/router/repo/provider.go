package repo

import (
	"context"
	"os"
	"svix-poc/package/app"

	"github.com/jackc/pgx/v4"
)

type Provider struct {
	*app.BaseActor
	conn *pgx.Conn
	Querier
}

func Provide() *Provider {
	return &Provider{}
}

func (p *Provider) Start(ctx context.Context) (first bool, err error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if first, err = p.BaseStart(ctx); first || err != nil {
		return first, err
	} else if p.conn, err = pgx.Connect(ctx, string(ProvideDbUrl())); err != nil {
		return true, nil
	} else {
		p.Querier = NewQuerier(p.conn)
	}
	return true, nil
}

func (p *Provider) Stop(ctx context.Context) (bool, error) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if last, err := p.BaseStop(ctx); !last || err != nil {
		return last, err
	}
	return true, p.conn.Close(ctx)
}

type DbUrl string

var DefaultDbUrl DbUrl = ""

func ProvideDbUrl() DbUrl {
	if url := DbUrl(os.Getenv("WEBHOOK_DB_URL")); url != "" {
		return url
	}
	return DefaultDbUrl
}
