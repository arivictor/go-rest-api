package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	connection, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &Database{Client: connection}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
