package database

import (
	"context"
	"fmt"
	"github.com/22Fariz22/crm-estate/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"net"
	"net/url"
)

func NewPsqlDB(ctx context.Context, c *config.Config) (*sqlx.DB, error) {
	dsn := &url.URL{
		Scheme: "postgres",
		Host:   net.JoinHostPort(c.Postgres.Host, c.Postgres.Port),
		User:   url.UserPassword(c.Postgres.User, c.Postgres.Password),
		Path:   c.Postgres.DBName,
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")
	q.Add("connect_timeout", "10")
	q.Add("application_name", "yourapp")
	dsn.RawQuery = q.Encode()

	db, err := sqlx.ConnectContext(ctx, "pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	db.SetMaxOpenConns(c.DB.MaxOpenConns)
	db.SetMaxIdleConns(c.DB.MaxIdleConns)
	db.SetConnMaxLifetime(c.DB.ConnMaxLifetime)
	db.SetConnMaxIdleTime(c.DB.ConnMaxIdleTime)

	log.Printf("PostgreSQL connected: %s | Pool: %d/%d",
		c.Postgres.DBName, c.DB.MaxOpenConns, c.DB.MaxIdleConns)

	return db, nil
}

func CloseDB(db *sqlx.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		} else {
			log.Println("DB connection closed")
		}
	}
}
