package controllers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"hcpb-api/configs"
)

var session *pgxpool.Pool

func Init(ctx context.Context) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome",
		configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		slog.Error("Error parsing DB config", "error", err)
		return
	}

	config.MaxConns = configs.DB_MAX_CONNS
	config.MinConns = configs.DB_MIN_CONNS
	config.MaxConnLifetime = configs.DB_MAX_CONN_LIFETIME
	config.MaxConnIdleTime = configs.DB_MAX_CONN_IDLE_TIME
	config.HealthCheckPeriod = configs.DB_HEALTH_CHECK_PERIOD

	session, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		return
	}

	if err = session.Ping(ctx); err != nil {
		slog.Error("Error pinging database", "error", err)
		return
	}

	slog.Info("Connected to database")
}

func Session() *pgxpool.Pool {
	return session
}
