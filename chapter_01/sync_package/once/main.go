package main

import (
	"database/sql"
	"sync"
)

type Repository struct {
	db *sql.DB
	dbInit sync.Once
}

func (r *Repository) Dial() error {
	var err error

	r.dbInit.Do(func() {
		r.db, err = sql.Open("mysql", "user:password@/dbname")
	})

	return err
}