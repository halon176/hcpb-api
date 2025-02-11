package controllers

import (
	"context"
	"database/sql"
	"log/slog"

	"hcpb-api/configs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dbURL = "postgres://" + configs.DB_USER + ":" + configs.DB_PASS + "@" + configs.DB_HOST + ":" + configs.DB_PORT + "/" + configs.DB_NAME + "?sslmode=disable"

func Migrate(ctx context.Context) {
	// Open database connection
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
		return
	}
	defer dbConn.Close()

	// Initialize migrations
	migrationsPath := "file://migrations"
	migration, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		slog.Error("Failed to initialize migration", "error", err)
		return
	}

	// Apply migrations
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("Failed to apply migrations", "error", err)
		return
	}

	slog.Info("Migrations applied successfully")
}
