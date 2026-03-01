package controllers

import (
	"context"
	"log"
)

func GetStatistics(ctx context.Context) (data []byte, err error) {
	ctx, span := tracer.Start(ctx, "db.GetStatistics")
	defer span.End()

	stmt := `
	SELECT json_build_object(
		'general_statistics', (SELECT json_agg(c) FROM (SELECT * FROM get_statistics()) c),
		'call_statistics', (SELECT json_agg(c) FROM (SELECT * FROM get_call_statistics()) c)
	) AS statistics
	`
	err = session.QueryRow(ctx, stmt).Scan(&data)
	if err != nil {
		span.RecordError(err)
		log.Println(err)
		return nil, err
	}
	return data, nil
}
