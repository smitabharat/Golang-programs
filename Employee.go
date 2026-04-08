package main

import "fmt"

type Employee map[string]interface{}

func main() {
	employees := []Employee{
		{"name": "smita", "Department": "IT", "Salary": 80000},
		{"name": "Revti", "Department": "IT", "Salary": 30000},
		{"name": "suraj", "Department": "IT", "Salary": 75000},
		{"name": "Mandar", "Department": "IT", "Salary": 60000},
	}
	department := "IT"
	totalSalary := CalculateSalary(employees, department)
	fmt.Printf("Total salary for the %s department is: %d\n", department, totalSalary)
}
func CalculateSalary(employees []Employee, Department string) int {
	totalSalary := 0
	for _, employee := range employees {
		if empDept, ok := employee["Department"].(string); ok && empDept == Department {
			if salary, ok := employee["Salary"].(int); ok {
				totalSalary += salary
			}
		}

	}
	return totalSalary
}
