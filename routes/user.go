package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brandtnick/rest/models"
)

// CreateUser - Handle create user route
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE USER")
	user := &models.User{}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := user.Create()
	json.NewEncoder(w).Encode(res)
}

// AuthUser - Handle user authentication
var AuthUser = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AUTH USER")
	user := &models.User{}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.Login(user.Username, user.Password)
	json.NewEncoder(w).Encode(res)
}

// GetUsers - Handle list of users
var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET USERS")

	res := models.GetUsers()
	json.NewEncoder(w).Encode(res)
}
