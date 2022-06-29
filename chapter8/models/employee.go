package models

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Position int

const (
	Undetermined Position = iota
	Junior
	Senior
	Manager
	CEO
)

type Employee struct {
	ID          int
	FullName    string
	Position    Position
	Salary      float64
	Joined      time.Time
	OnProbation int
	CreatedAt   time.Time
}

type Repository interface {
	Employees(ctx context.Context, pos Position) ([]Employee, error)
	Employee(ctx context.Context, ID int) (*Employee, error)
	Save(ctx context.Context, e *Employee) error
	Update(ctx context.Context, e *Employee) error
}
