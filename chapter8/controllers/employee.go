package controllers

import (
	"encoding/json"
	"fmt"
	"go_training/chapter8/models"
	"go_training/chapter8/services"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	employeeService *services.EmployeeService
}

func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

func (h *EmployeeController) GetEmployeeList(w http.ResponseWriter, r *http.Request) {
	var employeeList []models.Employee
	var err error
	query := r.URL.Query()
	posStr := query.Get("position")
	idPos, _ := strconv.Atoi(posStr)
	employeeList, err = h.employeeService.GetEmployees(idPos)
	fmt.Println(employeeList)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(employeeList)
}

func (h *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	var employee *models.Employee
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	employee, err = h.employeeService.GetSingleEmployee(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeToCreate models.Employee
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &employeeToCreate)
	created, err := h.employeeService.InsertEmployee(&employeeToCreate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if created {
		fmt.Println("Saved Employee Successfully")
	}
	json.NewEncoder(w).Encode(employeeToCreate)
}

func (h *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeToUpdate models.Employee
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &employeeToUpdate)
	updated, err := h.employeeService.UpdateEmployee(&employeeToUpdate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if updated {
		fmt.Println("Updated Employee Successfully")
	}
	json.NewEncoder(w).Encode(employeeToUpdate)
}
