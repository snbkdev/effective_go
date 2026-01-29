package main

import (
	"context"
	"database/sql"
	"sync"
)

var (
	cache = map[int]Country{}
	cacheMutex = &sync.Mutex{}
	db *sql.DB
)

type Country struct {
	Name string
	Population int
}

func UpdateCountryDyIDV1(ctx context.Context, countryID int) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	query := "select name, population from countries where id = $1"
	result := db.QueryRowContext(ctx, query, countryID)

	country := Country{}
	err := result.Scan(&country.Name, &country.Population)
	if err != nil {
		return err
	}

	cache[countryID] = country
	return nil
}
