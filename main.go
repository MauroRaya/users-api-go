package main

import (
	"encoding/json"
	"example/users-api-go/domain"
	"net/http"
)

var users []domain.User = []domain.User{
	{ID: 0, Name: "João"},
	{ID: 1, Name: "Maria"},
}

func main() {
	http.HandleFunc("/users", HandleGetUsers)
	http.HandleFunc("/users", HandleGetUser)

	http.ListenAndServe(":80", nil)
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := r.URL.Query().Get("id")

	json.NewEncoder(w).Encode(user)
}
