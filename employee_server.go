package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		employee := []Employee{
			{ID: "1", Name: "Vijay", Email: "vijay@mail.com", Balance: 10},
			{ID: "2", Name: "Kumar", Email: "kumar@mail.com", Balance: 20},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)
	})

	http.ListenAndServe(":8080", nil)
}