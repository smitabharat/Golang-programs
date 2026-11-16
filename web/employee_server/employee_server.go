// Simple REST API using the standard net/http package.
// GET http://localhost:8080/employee returns a list of employees as JSON.
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

var employees = []Employee{
	{ID: "1", Name: "Vijay", Email: "vijay@mail.com", Balance: 10},
	{ID: "2", Name: "Kumar", Email: "kumar@mail.com", Balance: 20},
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/employee", employeeHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
