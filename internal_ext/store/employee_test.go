package store

import (
	"strconv"
	"testing"

	"employee-management/internal_ext/models"
)

func TestCreateEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := store.CreateEmployee(emp)

	if createdEmp.ID == 0 {
		t.Errorf("expected ID to be set, got %d", createdEmp.ID)
	}
	if createdEmp.Name != emp.Name {
		t.Errorf("expected Name %s, got %s", emp.Name, createdEmp.Name)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	store := NewEmployeeStore()
	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := store.CreateEmployee(emp)

	retrievedEmp, err := store.GetEmployeeByID(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if retrievedEmp.Name != emp.Name {
		t.Errorf("expected Name %s, got %s", emp.Name, retrievedEmp.Name)
	}
}

func TestUpdateEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := store.CreateEmployee(emp)

	updateEmp := models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000}
	err := store.UpdateEmployee(createdEmp.ID, updateEmp)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	retrievedEmp, err := store.GetEmployeeByID(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if retrievedEmp.Name != updateEmp.Name {
		t.Errorf("expected Name %s, got %s", updateEmp.Name, retrievedEmp.Name)
	}
}

func TestDeleteEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := store.CreateEmployee(emp)

	err := store.DeleteEmployee(createdEmp.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = store.GetEmployeeByID(createdEmp.ID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestListEmployees(t *testing.T) {
	store := NewEmployeeStore()
	for i := 0; i < 25; i++ {
		store.CreateEmployee(models.Employee{Name: "Employee" + strconv.Itoa(i), Position: "Position", Salary: 50000})
	}

	employees, _ := store.ListEmployees(1, 10)
	if len(employees) != 10 {
		t.Errorf("expected 10 employees, got %d", len(employees))
	}

	employees, _ = store.ListEmployees(3, 10)
	if len(employees) != 5 {
		t.Errorf("expected 5 employees, got %d", len(employees))
	}
}
