package repositories

import (
	"go_training/chapter8/models"
	"go_training/chapter8/repositories"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	joined := time.Date(
		2022, 06, 07, 00, 00, 00, 00000000, time.UTC)
	created := time.Now()

	employeesTest := []models.Employee{
		{
			ID:          1,
			FullName:    "Luis Ramirez",
			Position:    2,
			Salary:      10,
			Joined:      joined,
			OnProbation: 1,
			CreatedAt:   created,
		},
		{
			ID:          2,
			FullName:    "Manuel Gutierrez",
			Position:    2,
			Salary:      20,
			Joined:      joined,
			OnProbation: 1,
			CreatedAt:   created,
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM employee WHERE position = ?")).
		WithArgs(2).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"}).
				AddRow(1, "Luis Ramirez", 2, 10, joined, 1, created).
				AddRow(2, "Manuel Gutierrez", 2, 20, joined, 1, created))

	employeeRepo := repositories.NewEmployeeRepo(db)
	employees, err := employeeRepo.Employees(2)
	assert.NoError(t, err)
	assert.Equal(t, employees, employeesTest)
}

func TestEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	joined := time.Date(
		2022, 06, 07, 00, 00, 00, 00000000, time.UTC)
	created := time.Now()

	employeeTest := &models.Employee{
		ID:          1,
		FullName:    "Luis Ramirez",
		Position:    2,
		Salary:      10,
		Joined:      joined,
		OnProbation: 1,
		CreatedAt:   created,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM employee WHERE id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"},
			).AddRow(
				1, "Luis Ramirez", 2, 10, joined, 1, created,
			),
		)

	employeeRepo := repositories.NewEmployeeRepo(db)
	employee, err := employeeRepo.Employee(1)
	assert.NoError(t, err)
	assert.Equal(t, employee, employeeTest)
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	joined := time.Date(
		2022, 06, 07, 00, 00, 00, 00000000, time.UTC)

	employeeTest := &models.Employee{
		ID:          1,
		FullName:    "Luis Ramirez",
		Position:    2,
		Salary:      10,
		Joined:      joined,
		OnProbation: 1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employee (id, full_name, position, salary, joined, on_probation) VALUES (?,?,?,?,?,?)")).
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	employeeRepo := repositories.NewEmployeeRepo(db)
	result, err := employeeRepo.Save(employeeTest)
	assert.NoError(t, err)
	assert.Equal(t, result, true)
}

func TestSaveUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	joined := time.Date(
		2022, 06, 07, 00, 00, 00, 00000000, time.UTC)

	employeeTest := &models.Employee{
		ID:          1,
		FullName:    "Luis Ramirez",
		Position:    2,
		Salary:      10,
		Joined:      joined,
		OnProbation: 1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employee SET (full_name, position, salary, joined, on_probation) VALUES (?,?,?,?,?) WHERE id = ?")).
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	employeeRepo := repositories.NewEmployeeRepo(db)
	result, err := employeeRepo.Update(employeeTest)
	assert.NoError(t, err)
	assert.Equal(t, result, true)
}
