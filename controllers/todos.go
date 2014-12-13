package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wless1/todos/models"
)

type TodosController struct{}

func (c *TodosController) deserializeTodo(r io.Reader) *models.Todo {
	decoder := json.NewDecoder(r)
	var t models.Todo

	if err := decoder.Decode(&t); err != nil {
		log.Fatal(err)
	}

	return &t
}

func (c *TodosController) Show(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	if t := models.GetTodo(id); t != nil {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(t)
	} else {
		fmt.Fprintln(res, "Not Found")
	}
}

func (c *TodosController) Delete(res http.ResponseWriter, req *http.Request) {
	result := models.DeleteTodo(mux.Vars(req)["id"])

	fmt.Fprintln(res, result)
}

func (c *TodosController) Update(res http.ResponseWriter, req *http.Request) {
	if t := c.deserializeTodo(req.Body); t != nil {
		id, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			log.Fatal(err)
		}
		t.Id = int64(id)

		models.UpdateTodo(t)

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(t)
	}
}

func (c *TodosController) Index(res http.ResponseWriter, req *http.Request) {
	todos := models.GetTodos()

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todos)
}

func (c *TodosController) Create(res http.ResponseWriter, req *http.Request) {
	if t := c.deserializeTodo(req.Body); t != nil {
		models.CreateTodo(t)

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(t)
	}
}
