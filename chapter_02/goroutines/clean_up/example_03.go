package main

import (
	"context"
	"database/sql"
	"time"
)

var db *sql.DB

func Example03() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	go func() {
		query := "select name, population from countries"
		results, err := db.QueryContext(ctx, query)
		if err != nil {
			return
		}

		userResults(results)
	}()
}

func userResults(rows *sql.Rows) {}