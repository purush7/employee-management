package services

import (
	"strconv"
	"testing"

	"employee-management/internal_ext/models"
	"employee-management/internal_ext/store"
)

func TestEmployeeService_CreateEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()
	service := NewEmployeeService(empStore)

	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := service.CreateEmployee(emp)

	if createdEmp.ID == 0 {
		t.Errorf("expected ID to be set, got %d", createdEmp.ID)
	}
	if createdEmp.Name != emp.Name {
		t.Errorf("expected Name %s, got %s", emp.Name, createdEmp.Name)
	}
}

func TestEmployeeService_GetEmployeeByID(t *testing.T) {
	empStore := store.NewEmployeeStore()
	service := NewEmployeeService(empStore)

	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := service.CreateEmployee(emp)

	retrievedEmp, err := service.GetEmployeeByID(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if retrievedEmp.Name != emp.Name {
		t.Errorf("expected Name %s, got %s", emp.Name, retrievedEmp.Name)
	}
}

func TestEmployeeService_UpdateEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()
	service := NewEmployeeService(empStore)

	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := service.CreateEmployee(emp)

	updateEmp := models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000}
	err := service.UpdateEmployee(createdEmp.ID, updateEmp)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	retrievedEmp, err := service.GetEmployeeByID(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if retrievedEmp.Name != updateEmp.Name {
		t.Errorf("expected Name %s, got %s", updateEmp.Name, retrievedEmp.Name)
	}
}

func TestEmployeeService_DeleteEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()
	service := NewEmployeeService(empStore)

	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := service.CreateEmployee(emp)

	err := service.DeleteEmployee(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = service.GetEmployeeByID(createdEmp.ID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestEmployeeService_ListEmployees(t *testing.T) {
	empStore := store.NewEmployeeStore()
	service := NewEmployeeService(empStore)

	for i := 0; i < 25; i++ {
		service.CreateEmployee(models.Employee{Name: "Employee" + strconv.Itoa(i), Position: "Position", Salary: 50000})
	}

	employees, _ := service.ListEmployees(1, 10)
	if len(employees) != 10 {
		t.Errorf("expected 10 employees, got %d", len(employees))
	}

	employees, _ = service.ListEmployees(3, 10)
	if len(employees) != 5 {
		t.Errorf("expected 5 employees, got %d", len(employees))
	}
}
