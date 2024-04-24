package main

import (
	"encoding/json"
	"net/http"

	"github.com/rafack/digport-academy/model"
)

func StartServer() {
	http.HandleFunc("/employees", employeesHandler)

	http.ListenAndServe(":8080", nil)
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getEmployees(w, r)
	} else if r.Method == "POST" {
		addEmployee(w, r)
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	registerEmployee(employee)

	w.WriteHeader(http.StatusCreated)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		employeesFiltradosPorNome := employeesByName(queryNome)
		json.NewEncoder(w).Encode(employeesFiltradosPorNome)
	} else {
		employees := Employees
		json.NewEncoder(w).Encode(employees)
	}
}
