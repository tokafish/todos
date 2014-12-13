package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	c := &TodosController{}

	router := mux.NewRouter()
	todos := router.Headers("Content-Type", "application/json").Subrouter()

	todos.Methods("GET").Path("/todos").HandlerFunc(c.Index)
	todos.Methods("POST").Path("/todos").HandlerFunc(c.Create)
	todos.Methods("GET").Path("/todos/{id:[0-9]+}").HandlerFunc(c.Show)
	todos.Methods("DELETE").Path("/todos/{id:[0-9]+}").HandlerFunc(c.Delete)
	todos.Methods("PATCH").Path("/todos/{id:[0-9]+}").HandlerFunc(c.Update)

	return router
}
