package repo

import (
	"context"
	"os"
	"svix-poc/lib/app"

	"github.com/jackc/pgx/v4"
)

const SingletonKey = "router_repo"

type Repository struct {
	*app.BaseActor
	conn *pgx.Conn
	Querier
}

func Provide() *Repository {
	return &Repository{
		BaseActor: app.NewActor(nil),
	}
}

func (p *Repository) Start(ctx context.Context) (first bool, err error) {
	if first, err = p.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	} else if p.conn, err = pgx.Connect(ctx, string(ProvideDbUrl())); err != nil {
		return true, nil
	} else {
		p.Querier = NewQuerier(p.conn)
	}
	return true, nil
}

func (p *Repository) Stop(ctx context.Context) (bool, error) {
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
	if url := DbUrl(os.Getenv("WEBHOOK_DB_URL")); url != "" {
		return url
	}
	return DefaultDbUrl
}
