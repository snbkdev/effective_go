package main

import "context"

func UpdateCountryByIDV2(ctx context.Context, countryID int) error {
	query := "select name, population from countries where id = $1"
	result := db.QueryRowContext(ctx, query, countryID)

	country := Country{}
	err := result.Scan(&country.Name, &country.Population)
	if err != nil {
		return err
	}

	cacheMutex.Lock()
	cache[countryID] = country
	cacheMutex.Unlock()

	return nil
}