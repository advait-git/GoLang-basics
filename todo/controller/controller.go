package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/advait-git/todolist/models"
	"github.com/advait-git/todolist/service"
	"github.com/gorilla/mux"
)

// create todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	//define todo model
	var todo models.Todo
	//decode the request body into the todo model
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//call the service layer to create the todo
	createdTodo, err := service.CreateTodo(todo)
	//return the response
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)

}

// get all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	//call GetAllTodos service
	todos, err := service.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to get the todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

//delete todo

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//get the id from req
	params := mux.Vars(r)
	id := params["id"]
	//call the service layer to delete the todo
	err := service.DeleteTodo(id)
	if err != nil {
		log.Fatal("Error in delete api controller", err)
	}
	json.NewEncoder(w).Encode("Todo deleted successfully")
}

// update todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	
	params :=mux.Vars(r)
	id := params["id"]

	var updatedData models.Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedData)

	err := service.UpdateTodo(id , updatedData)
	if err != nil {
		log.Fatal("Error in update api controller", err)
	}
	json.NewEncoder(w).Encode("Todo Updation Completed")
}
