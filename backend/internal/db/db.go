package db

import (
	"context"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
		cfg.DBSSLMode)

	fmt.Println(connectionString)

	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to database, %w", err)
	}

	return &Database{
		Client: dbConn,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
