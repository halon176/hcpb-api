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
	var err error

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome",
		configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT)

	session, err = pgxpool.New(ctx, connStr)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
	}

	err = session.Ping(ctx)
	if err != nil {
		slog.Error("Error pinging database", "error", err)
	}

	slog.Info("Connected to database")

}
