package main

import (
	"api_postgres/configs"
	"api_postgres/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/list", handlers.List)
	r.Get("/{id}", handlers.Get)
	println(configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
