package controllers

import (
	"context"
	"fmt"
	s "hcpb-api/schemas"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("hcpb-api/controllers")

func GetLastCalls(ctx context.Context) (data []byte, err error) {
	ctx, span := tracer.Start(ctx, "db.GetLastCalls")
	defer span.End()

	stmt := `
	SELECT json_agg(c) FROM (
		SELECT services.name as service_name, types.name as type_name, calls.chat_id, calls.coin, calls.created_at
		FROM calls
		JOIN services ON calls.service_id = services.id
		JOIN types ON calls.type_id = types.id
		ORDER BY calls.created_at DESC
		LIMIT 10
	) c
	`
	err = session.QueryRow(ctx, stmt).Scan(&data)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error getting last calls", "error", err)
		return nil, err
	}
	return data, nil
}

func GetPaginatedCalls(ctx context.Context, params s.PaginationParams) (data []byte, err error) {
	ctx, span := tracer.Start(ctx, "db.GetPaginatedCalls")
	defer span.End()

	span.SetAttributes(
		attribute.Int("db.page", params.Page),
		attribute.Int("db.page_size", params.PageSize),
	)

	stmt := fmt.Sprintf(`
	SELECT json_agg(c) FROM (
		SELECT services.name as service_name, types.name as type_name, calls.chat_id, calls.coin, calls.created_at
		FROM calls
		JOIN services ON calls.service_id = services.id
		JOIN types ON calls.type_id = types.id
		ORDER BY calls.created_at DESC
		LIMIT %d OFFSET %d
	) c
	`, params.PageSize, params.Offset)

	err = session.QueryRow(ctx, stmt).Scan(&data)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error getting paginated calls", "error", err)
		return nil, err
	}
	return data, nil
}

func GetTotalCallsCount(ctx context.Context) (int64, error) {
	ctx, span := tracer.Start(ctx, "db.GetTotalCallsCount")
	defer span.End()

	var count int64
	stmt := `SELECT COUNT(*) FROM calls`
	err := session.QueryRow(ctx, stmt).Scan(&count)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error getting total calls count", "error", err)
		return 0, err
	}
	return count, nil
}

func InsertCall(ctx context.Context, call s.Call) error {
	ctx, span := tracer.Start(ctx, "db.InsertCall")
	defer span.End()

	span.SetAttributes(
		attribute.Int("call.service_id", call.ServiceID),
		attribute.Int("call.type_id", call.TypeID),
		attribute.String("call.coin", call.Coin),
	)

	stmt := `
	INSERT INTO calls (service_id, type_id, chat_id, coin) VALUES ($1, $2, $3, $4)
	`
	_, err := session.Exec(ctx, stmt, call.ServiceID, call.TypeID, call.ChatID, call.Coin)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error inserting call", "error", err)
		return err
	}
	return nil
}
