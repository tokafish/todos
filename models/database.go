package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseSession struct {
	*sqlx.DB
	databaseName string
}

var db *DatabaseSession = connect("todos_dev")

func connect(name string) *DatabaseSession {
	db := sqlx.MustConnect("postgres", "user=postgres dbname="+name+" sslmode=disable")

	return &DatabaseSession{db, name}
}
