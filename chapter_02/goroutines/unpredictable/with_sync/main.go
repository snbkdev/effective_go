package main

import (
	"database/sql"
	"fmt"
	"sync"
)

func NewLoader() *Loader {
	loader := &Loader{
		dbInit: &sync.Cond{
			L: &sync.Mutex{},
		},
	}

	go loader.initializeDBPool()

	return loader
}

type Loader struct {
	db *sql.DB

	dbInit *sync.Cond
}

func(l *Loader) LoadAll() error {
	l.dbInit.L.Lock()
	for l.db == nil {
		l.dbInit.Wait()
	}

	l.dbInit.L.Unlock()

	results, err := l.db.Query("select user from users")
	if err != nil {
		return err
	}

	var username string
	for results.Next() {
		err := results.Scan(&username)
		if err != nil {
			return err
		}

		fmt.Printf("User: %#v\n", username)
	}

	return nil
}

func(l *Loader) initializeDBPool() {
	defer l.dbInit.Broadcast()

	var err error
	l.db, err = sql.Open("mysql", "root@tcp(localhost:3306)/mysql?autocommit=true")
	if err != nil {
		panic(err)
	}
}