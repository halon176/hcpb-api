package controllers

import (
	"context"
)

func GetExcluded(ctx context.Context) ([]byte, error) {
	stmt := `
	SELECT json_agg(item) FROM excluded
	`
	var jsonData []byte
	err := session.QueryRow(ctx, stmt).Scan(&jsonData)
	if err != nil {
		return nil, err
	}

	if string(jsonData) == "" {
		jsonData = []byte("[]")
	}

	return jsonData, nil
}

func InsertExcluded(ctx context.Context, item string) error {
	stmt := `
	INSERT INTO excluded (item) VALUES ($1)
	`
	_, err := session.Exec(ctx, stmt, item)
	if err != nil {
		return err
	}

	return nil
}

func DeleteExcluded(ctx context.Context, item string) error {
	stmt := `
	DELETE FROM excluded WHERE item = $1
	`
	_, err := session.Exec(ctx, stmt, item)
	if err != nil {
		return err
	}

	return nil
}
