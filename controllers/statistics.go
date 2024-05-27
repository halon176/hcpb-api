package controllers

import (
	"context"
	"log"
)

func GetStatistics(ctx context.Context) (data []byte, err error) {
	stmt := `
	SELECT json_agg(c) from (SELECT * from get_statistics()) c
	`
	err = session.QueryRow(ctx, stmt).Scan(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}
