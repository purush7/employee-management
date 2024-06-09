package services

import (
	"employee-management/internal_ext/models"
	"employee-management/internal_ext/store"
)

type EmployeeService struct {
	store *store.EmployeeStore
}

func NewEmployeeService(store *store.EmployeeStore) *EmployeeService {
	return &EmployeeService{store: store}
}

func (s *EmployeeService) CreateEmployee(emp models.Employee) models.Employee {
	return s.store.CreateEmployee(emp)
}

func (s *EmployeeService) GetEmployeeByID(id int) (models.Employee, error) {
	return s.store.GetEmployeeByID(id)
}

func (s *EmployeeService) UpdateEmployee(id int, emp models.Employee) error {
	return s.store.UpdateEmployee(id, emp)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.store.DeleteEmployee(id)
}

func (s *EmployeeService) ListEmployees(page, pageSize int) ([]models.Employee,int) {
	return s.store.ListEmployees(page, pageSize)
}
