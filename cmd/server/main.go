package main

import (
	"log"
	"net/http"

	"employee-management/internal_ext/controllers"
	"employee-management/internal_ext/routes"
	"employee-management/internal_ext/services"
	"employee-management/internal_ext/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	store := store.NewEmployeeStore()
	service := services.NewEmployeeService(store)
	controller := controllers.NewEmployeeController(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routes.NewRouter(controller))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
