package main

import (
	"net/http"
	"os"

	"github.com/wless1/todos/controllers"
)

func main() {
	http.Handle("/", controllers.Routes())
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
