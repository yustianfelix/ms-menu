package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MenuItem represents the data structure for a menu item
type MenuItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory data store
var menu = []MenuItem{
	{ID: "1", Name: "Coffee", Price: 35000},
	{ID: "2", Name: "Tea", Price: 15000},
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

func main() {
	// Define routes
	http.HandleFunc("/menu", getMenu)

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
