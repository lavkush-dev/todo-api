package main

import (
	"log"
	"net/http"
	"todo-api/db"
	"todo-api/handlers"
	"todo-api/middleware"
	"todo-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	r := mux.NewRouter()

	// Public route
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// mock login
		token, _ := utils.GenerateJWT("testuser")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token":"` + token + `"}`))
	}).Methods("POST")

	// Protected route
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)
	api.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	api.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
	api.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
