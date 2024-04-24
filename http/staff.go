package main

import "github.com/rafack/digport-academy/model"

var Employees []model.Employee = []model.Employee{}

func employeesByName(name string) []model.Employee {
	result := []model.Employee{}

	for _, employee := range Employees {
		if employee.Name == name {
			result = append(result, employee)
		}
	}

	return result
}

func registerEmployee(employee model.Employee) {
	Employees = append(Employees, employee)
}
