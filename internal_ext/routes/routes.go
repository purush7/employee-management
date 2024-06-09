package routes

import (
	"employee-management/internal_ext/controllers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controller *controllers.EmployeeController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/employee", controller.CreateEmployee)
	r.Get("/employee/{id}", controller.GetEmployee)
	r.Put("/employee/{id}", controller.UpdateEmployee)
	r.Delete("/employee/{id}", controller.DeleteEmployee)
	r.Get("/employee", controller.ListEmployees)

	return r
}
