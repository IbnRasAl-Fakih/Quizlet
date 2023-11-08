package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	connection *sql.DB
}

var instance *Database

func GetDatabaseInstance() *Database {
	if instance == nil {
		connStr := "user=postgres password=55810579 dbname=quizlet sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		instance = &Database{connection: db}
	}

	return instance
}
