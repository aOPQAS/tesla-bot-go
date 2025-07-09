package pgsql

import (
	"errors"
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/tesla/tesla-bot-go/config"
)

const (
	migrationsDir = "migrations"
)

type Client struct {
	conn *dbr.Connection
}

func NewCLient(conn *dbr.Connection) *Client {
	return &Client{conn: conn}
}

func (c *Client) GetSession() *dbr.Session {
	return c.conn.NewSession(nil)
}

func (c *Client) RunMigrations() error {
	driver, err := postgres.WithInstance(c.conn.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Failed to create driver: %w", err)
	}

	migrationsPath := fmt.Sprintf("file://%s", migrationsDir)
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}

func CreatePostgresConnection(cfg config.PostgresConfig) (*dbr.Connection, error) {
	cs := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)

	conn, err := dbr.Open("postgres", cs, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping connection: %w", err)
	}

	return conn, nil
}
