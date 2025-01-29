package api

import (
	"github.com/gorilla/mux"
	"github.com/advait-git/todolist/controller"
)


func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/todo", controller.CreateTodo).Methods("POST")
	r.HandleFunc("/api/v1/todo/{id}", controller.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/api/v1/todos", controller.GetAllTodos).Methods("GET")
	r.HandleFunc("/api/v1/todo/{id}", controller.UpdateTodo).Methods("PUT")

	return r
}