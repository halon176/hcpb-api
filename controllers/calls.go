package controllers

import (
	"context"
	"fmt"
	s "hcpb-api/schemas"
	"log/slog"
)

func GetLastCalls(ctx context.Context) (data []byte, err error) {
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
		slog.Error("Error getting last calls", "error", err)
		return nil, err
	}
	return data, nil
}

func GetPaginatedCalls(ctx context.Context, params s.PaginationParams) (data []byte, err error) {
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
		slog.Error("Error getting paginated calls", "error", err)
		return nil, err
	}
	return data, nil
}

func GetTotalCallsCount(ctx context.Context) (int64, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM calls`
	err := session.QueryRow(ctx, stmt).Scan(&count)
	if err != nil {
		slog.Error("Error getting total calls count", "error", err)
		return 0, err
	}
	return count, nil
}

func InsertCall(ctx context.Context, call s.Call) error {
	stmt := `
	INSERT INTO calls (service_id, type_id, chat_id, coin) VALUES ($1, $2, $3, $4)
	`
	_, err := session.Exec(ctx, stmt, call.ServiceID, call.TypeID, call.ChatID, call.Coin)
	if err != nil {
		slog.Error("Error inserting call", "error", err)
		return err
	}
	return nil
}
