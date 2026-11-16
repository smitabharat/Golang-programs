// Program to calculate the total salary of all employees in a department.
package main

import "fmt"

// Employee uses a struct instead of map[string]interface{} so the
// compiler checks field names and types for us.
type Employee struct {
	Name       string
	Department string
	Salary     int
}

func CalculateSalary(employees []Employee, department string) int {
	total := 0
	for _, emp := range employees {
		if emp.Department == department {
			total += emp.Salary
		}
	}
	return total
}

func main() {
	employees := []Employee{
		{Name: "Smita", Department: "IT", Salary: 80000},
		{Name: "Revti", Department: "IT", Salary: 30000},
		{Name: "Suraj", Department: "IT", Salary: 75000},
		{Name: "Mandar", Department: "HR", Salary: 60000},
	}

	for _, dept := range []string{"IT", "HR"} {
		fmt.Printf("Total salary for the %s department is: %d\n", dept, CalculateSalary(employees, dept))
	}
}

// Output:
// Total salary for the IT department is: 185000
// Total salary for the HR department is: 60000
