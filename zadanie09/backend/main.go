package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Payment struct {
	Amount float64 `json:"amount"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodGet {
		products := []Product{
			{ID: 1, Name: "Test_01", Price: 22.22},
			{ID: 2, Name: "Test_02", Price: 33.33},
			{ID: 3, Name: "Test_03", Price: 44.44},
		}
		json.NewEncoder(w).Encode(products)
	}
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var p Payment
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Received payment: %.2f PLN\n", p.Amount)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/payments", paymentsHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
