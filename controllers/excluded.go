package controllers

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

func GetExcluded(ctx context.Context) ([]byte, error) {
	ctx, span := tracer.Start(ctx, "db.GetExcluded")
	defer span.End()

	stmt := `
	SELECT json_agg(item) FROM excluded
	`
	var jsonData []byte
	err := session.QueryRow(ctx, stmt).Scan(&jsonData)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if string(jsonData) == "" {
		jsonData = []byte("[]")
	}

	return jsonData, nil
}

func InsertExcluded(ctx context.Context, item string) error {
	ctx, span := tracer.Start(ctx, "db.InsertExcluded")
	defer span.End()

	span.SetAttributes(attribute.String("excluded.item", item))

	stmt := `
	INSERT INTO excluded (item) VALUES ($1)
	`
	_, err := session.Exec(ctx, stmt, item)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

func DeleteExcluded(ctx context.Context, item string) error {
	ctx, span := tracer.Start(ctx, "db.DeleteExcluded")
	defer span.End()

	span.SetAttributes(attribute.String("excluded.item", item))

	stmt := `
	DELETE FROM excluded WHERE item = $1
	`
	_, err := session.Exec(ctx, stmt, item)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
