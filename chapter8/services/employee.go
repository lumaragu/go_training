package services

import (
	"go_training/chapter8/models"
	"go_training/chapter8/repositories"
)

type EmployeeService struct {
	employeeRepo *repositories.EmployeeRepo
}

func NewEmployeeService(employeeRepo *repositories.EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}

func (s *EmployeeService) ExistEmployee(id int) (bool, error) {
	employee, err := s.employeeRepo.Employee(id)
	if err == nil && employee != nil {
		return true, nil
	}
	return false, err
}

func (s *EmployeeService) GetEmployees(pos int) ([]models.Employee, error) {
	employeeList, err := s.employeeRepo.Employees(pos)
	return employeeList, err
}

func (s *EmployeeService) GetSingleEmployee(ID int) (*models.Employee, error) {
	item, err := s.employeeRepo.Employee(ID)
	return item, err
}

func (s *EmployeeService) InsertEmployee(employee *models.Employee) (bool, error) {
	state, err := s.employeeRepo.Save(employee)
	return state, err
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) (bool, error) {
	var err error
	found, err := s.ExistEmployee(employee.ID)
	if !found {
		return false, err
	}
	state, err := s.employeeRepo.Update(employee)
	return state, err
}
