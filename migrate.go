package database

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func MigrateDB(dsn string) error {
	m, err := migrate.New("embed://migrations", dsn)
	if err != nil {
		return fmt.Errorf("create migrate: %w", err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migrations to apply")
			return nil
		}
		return fmt.Errorf("migrate up: %w", err)
	}
	log.Println("Migrations applied successfully")
	return nil
}
