package db

import "os"
import "path/filepath"

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("migrating our database")

	execPath, _ := os.Executable()
	migrationPath := filepath.Join(filepath.Dir(execPath), "migrations")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations: %w", err)
		}
	}

	// Force migration version if dirty
	err = m.Force(1) // Replace with the version you need to force
	if err != nil {
		log.Fatalf("Could not force migration version: %v", err)
	}

	fmt.Println("successfully migrated the database")

	return nil
}
