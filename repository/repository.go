package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const userName string = "auli"
const password string = "bq47w8mz"
const dbName string = "test"
const host string = "127.0.0.1"
const port string = "5432"

type PostgreRSSRepository struct {
	pool *pgxpool.Pool
}

func NewPostgreRSSRepository() (*PostgreRSSRepository, error) {
	DBUri := "postgresql://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbName
	config, err := pgxpool.ParseConfig(DBUri)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create config")
	}

	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create pool")
	}
	return &PostgreRSSRepository{pool: pool}, nil
}
