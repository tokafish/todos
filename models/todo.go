package models

import (
	"database/sql"
	"log"
)

type Todo struct {
	Id        int64
	Text      string
	Completed bool
}

func GetTodo(id interface{}) *Todo {
	var todo Todo
	err := db.Get(&todo, "SELECT * FROM todos WHERE id = $1 LIMIT 1", id)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		log.Fatal(err)
	}

	return &todo
}

func GetTodos() []Todo {
	todos := []Todo{}

	err := db.Select(&todos, "SELECT * FROM todos")

	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func CreateTodo(todo *Todo) {
	err := db.QueryRow("INSERT INTO todos (text, completed) VALUES ($1, $2) RETURNING id", todo.Text, todo.Completed).Scan(&todo.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTodo(todo *Todo) {
	_, err := db.Exec("UPDATE todos SET text = $1, completed = $2 WHERE id = $3", todo.Text, todo.Completed, todo.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodo(id interface{}) bool {
	result, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()

	return (err == nil && rows == 1)
}
