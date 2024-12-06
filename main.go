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
	r.HandleFunc("/users", HandleAddUser).Methods("POST")
	r.HandleFunc("/users/{id}", HandleEditUser).Methods("PATCH")

	http.ListenAndServe(":80", r)
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid parameter value"})
		return
	}

	for i := range users {
		if users[i].ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
}

func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newUser domain.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json data"})
		return
	}

	if newUser.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	for i := range users {
		if users[i].ID == newUser.ID {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "id already exists"})
			return
		}
	}

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func HandleEditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid parameter value"})
		return
	}

	var userDB *domain.User

	for i := range users {
		if users[i].ID == id {
			userDB = &users[i]
			break
		}
	}

	if userDB == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	var editedUser domain.User

	err = json.NewDecoder(r.Body).Decode(&editedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json data"})
		return
	}

	if editedUser.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	userDB.Name = editedUser.Name

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user edited successfully",
		"user":    *userDB},
	)
}
