package repositories

import (
	"database/sql"
	"fmt"
	"go_training/chapter8/models"
)

type EmployeeRepo struct {
	dbClient *sql.DB
}

func NewEmployeeRepo(dbClient *sql.DB) *EmployeeRepo {
	return &EmployeeRepo{
		dbClient: dbClient,
	}
}

func (e *EmployeeRepo) Employees(pos int) ([]models.Employee, error) {
	var err error
	rows, err := e.dbClient.Query("SELECT * FROM employee WHERE position = ?", pos)
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	employeeList := make([]models.Employee, 0)

	for rows.Next() {
		var employee models.Employee
		err = rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Position,
			&employee.Salary,
			&employee.Joined,
			&employee.OnProbation,
			&employee.CreatedAt,
		)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		employeeList = append(employeeList, employee)
	}
	return employeeList, nil
}

func (e *EmployeeRepo) Employee(ID int) (*models.Employee, error) {
	var err error
	var employee models.Employee
	err = e.dbClient.QueryRow("SELECT * FROM employee WHERE id = ?", ID).Scan(
		&employee.ID,
		&employee.FullName,
		&employee.Position,
		&employee.Salary,
		&employee.Joined,
		&employee.OnProbation,
		&employee.CreatedAt,
	)
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	return &employee, nil
}

func (e *EmployeeRepo) Save(employee *models.Employee) (bool, error) {
	var err error
	query, err := e.dbClient.Prepare("INSERT INTO employee (id, full_name, position, salary, joined, on_probation) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("ERROR INSERT QUERY - %s", err)
		return false, err
	}
	query.Exec(employee.ID, employee.FullName, employee.Position, employee.Salary, employee.Joined, employee.OnProbation)
	return true, nil
}

func (e *EmployeeRepo) Update(employee *models.Employee) (bool, error) {
	var err error
	query, err := e.dbClient.Prepare("UPDATE employee SET (full_name, position, salary, joined, on_probation) VALUES (?,?,?,?,?) WHERE id = ?")
	if err != nil {
		fmt.Printf("ERROR UPDATE QUERY - %s", err)
		return false, err
	}
	query.Exec(employee.FullName, employee.Position, employee.Salary, employee.Joined, employee.OnProbation, employee.ID)
	return true, nil
}
