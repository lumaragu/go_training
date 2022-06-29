package main

import (
	"fmt"
	"go_training/chapter8/controllers"
	"go_training/chapter8/db"
	"go_training/chapter8/repositories"
	"go_training/chapter8/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := db.ConnectDB()
	defer db.Close()
	employeeRepo := repositories.NewEmployeeRepo(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	router := mux.NewRouter()
	fmt.Println("Server is running on Port 3333")

	router.HandleFunc("/api/v1/employee", employeeController.GetEmployeeList).Methods("GET")
	router.HandleFunc("/api/v1/employee/{id}", employeeController.GetEmployee).Methods("GET")
	router.HandleFunc("/api/v1/employee", employeeController.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/v1/employee/{id}", employeeController.UpdateEmployee).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3333", router))
}
