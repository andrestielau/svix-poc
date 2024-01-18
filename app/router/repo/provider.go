package repo

import (
	"context"
	"os"
	"svix-poc/lib/app"

	"github.com/jackc/pgx/v4"
)

const SingletonKey = "router_repo"

type Repository struct {
	Querier
}
type Provider struct {
	*app.BaseActor
	conn *pgx.Conn
	*Repository
}

func Provide() *Provider {
	return &Provider{
		BaseActor:  app.NewActor(nil),
		Repository: &Repository{},
	}
}

func (p *Provider) Start(ctx context.Context) (first bool, err error) {
	if first, err = p.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	} else if p.conn, err = pgx.Connect(ctx, string(ProvideDbUrl())); err != nil {
		return true, err
	}
	p.Querier = NewQuerier(p.conn)
	return true, nil
}

func (p *Provider) Stop(ctx context.Context) (bool, error) {
	if last, err := p.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	if p.conn == nil {
		return true, nil
	}
	return true, p.conn.Close(ctx)
}

type DbUrl string

var DefaultDbUrl DbUrl = "postgres://postgres:postgres@localhost:5432?sslmode=disable"

func ProvideDbUrl() DbUrl {
	if url := DbUrl(os.Getenv("ROUTER_DB_URL")); url != "" {
		return url
	}
	return DefaultDbUrl
}
