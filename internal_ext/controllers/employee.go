package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"employee-management/internal_ext/models"
	"employee-management/internal_ext/services"
	"employee-management/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	createdEmployee := c.service.CreateEmployee(emp)
	utils.RespondWithJSON(w, http.StatusCreated, createdEmployee)
}

func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}
	emp, err := c.service.GetEmployeeByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, emp)
}

func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.UpdateEmployee(id, emp); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	type responseStruct struct {
		Message string `json:"message"`
	}
	utils.RespondWithJSON(w, http.StatusOK, responseStruct{
		Message: "success",
	})
}

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	if err := c.service.DeleteEmployee(id); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	type responseStruct struct {
		Message string `json:"message"`
	}
	utils.RespondWithJSON(w, http.StatusOK, responseStruct{
		Message: "success",
	})
}

func (c *EmployeeController) ListEmployees(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	employees, totalSize := c.service.ListEmployees(page, pageSize)
	type responseStruct struct {
		Employees []models.Employee
		TotalSize int
	}
	response := responseStruct{
		Employees: employees,
		TotalSize: totalSize,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
