package main

import (
	"encoding/json"
	"example/users-api-go/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users []domain.User = []domain.User{
	{ID: 0, Name: "João"},
	{ID: 1, Name: "Maria"},
}

func main() {
	fmt.Println("Servidor rodando na porta 80")
	r := mux.NewRouter()

	r.HandleFunc("/users", HandleGetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", HandleGetUser).Methods("GET")

	http.ListenAndServe(":80", r)
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid parameter value"})
		return
	}

	for i := range users {
		if users[i].ID == id {
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
}
