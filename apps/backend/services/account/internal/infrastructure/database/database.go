package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PostgresDB struct {
	connStr string
	Client *sql.DB
}

func NewPostgresDB(connStr string) (*PostgresDB, error) {
    client, err := sql.Open("postgres", connStr); if err != nil {
		return nil, err
	}

	return &PostgresDB{Client: client, connStr: connStr}, nil
}

func (db *PostgresDB)  RunMigrations() error {
	migrationsPath := "file://internal/infrastructure/database/migrations"

	m, err := migrate.New(migrationsPath, db.connStr)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
		return err
	}

	log.Println("Migrations ran successfully")
	return nil
}

func (db *PostgresDB)  Close() {
	db.Client.Close()
}