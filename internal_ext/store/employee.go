package store

import (
	"employee-management/internal_ext/models"
	"errors"
	"sync"
)

type EmployeeStore struct {
	mu        sync.Mutex
	employees map[int]models.Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
}

func (store *EmployeeStore) CreateEmployee(emp models.Employee) models.Employee {
	store.mu.Lock()
	defer store.mu.Unlock()

	emp.ID = store.nextID
	store.employees[emp.ID] = emp
	store.nextID++
	return emp
}

func (store *EmployeeStore) GetEmployeeByID(id int) (models.Employee, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	emp, exists := store.employees[id]
	if !exists {
		return models.Employee{}, errors.New("employee not found")
	}
	return emp, nil
}

func (store *EmployeeStore) UpdateEmployee(id int, emp models.Employee) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[id]; !exists {
		return errors.New("employee not found")
	}
	emp.ID = id
	store.employees[id] = emp
	return nil
}

func (store *EmployeeStore) DeleteEmployee(id int) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[id]; !exists {
		return errors.New("employee not found")
	}
	delete(store.employees, id)
	return nil
}

func (store *EmployeeStore) ListEmployees(page, pageSize int) ([]models.Employee, int) {
	store.mu.Lock()
	defer store.mu.Unlock()

	start := (page - 1) * pageSize
	end := start + pageSize

	var employees []models.Employee
	count := 0
	for id := 1; id < store.nextID; id++ {
		if emp, exists := store.employees[id]; exists {
			if count >= start && count < end {
				employees = append(employees, emp)
			}
			count++
		}
		if count >= end {
			break
		}
	}

	return employees, count
}
