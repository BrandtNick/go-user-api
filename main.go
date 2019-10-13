package main

import (
	"fmt"
	"net/http"
	"os"

	mw "github.com/brandtnick/rest/middlewares"
	routes "github.com/brandtnick/rest/routes"
	"github.com/fatih/color"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r)) // apply cors middleware to all routes
	r.Use(mw.LogHandler)

	// protected routes
	pr := r.PathPrefix("").Subrouter()
	pr.Use(mw.Authorization) // attach JWT authorization middleware to protected routes
	pr.HandleFunc("/users", routes.GetUsers).Methods("GET")

	// routes
	r.HandleFunc("/users", routes.CreateUser).Methods("POST")
	r.HandleFunc("/users/login", routes.AuthUser).Methods("POST")

	// routes not found
	r.HandleFunc("", routes.NotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	color.Green("listening on port: %s", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Print(err)
	}
}
