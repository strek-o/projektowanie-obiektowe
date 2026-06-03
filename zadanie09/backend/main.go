package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status   string `json:"status"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}

var users = map[string]string{
	"user":  "user",
	"admin": "admin",
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(LoginResponse{Status: "error", Message: "invalid request"})
		return
	}

	password, ok := users[req.Username]
	if !ok || password != req.Password {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(LoginResponse{Status: "error", Message: "Failure!"})
		return
	}

	_ = json.NewEncoder(w).Encode(LoginResponse{Status: "ok", Username: req.Username})
}

func main() {
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server: %v", err)
	}
}
