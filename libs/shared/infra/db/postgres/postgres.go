package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"log"
)

type Context struct {
	connected      bool
	pool           *pgxpool.Pool
	programContext context.Context
}

var pgContext *Context

func NewContext(ctx context.Context) (*Context, error) {
	if pgContext != nil {
		return pgContext, nil
	}

	pool, err := connect()
	if err != nil {
		return &Context{connected: false}, err
	}

	pgContext = &Context{true, pool, ctx}

	return pgContext, nil
}

func (ctx *Context) IsConnected() bool {
	return ctx.connected
}

func (ctx *Context) Acquire() (*pgxpool.Conn, error) {
	conn, err := ctx.pool.Acquire(ctx.programContext)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to acquire the connection from %s", "target")
	}

	return conn, nil
}

func connect() (*pgxpool.Pool, error) {
	connectionString := "postgres://appadmin:@pp@m1n@localhost:15342/gotamboon_local"
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err.Error())

		return nil, err
	}

	return pool, nil
}
