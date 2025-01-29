package service

import (
	"errors"

	"github.com/advait-git/todolist/models"
	"github.com/advait-git/todolist/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(todo models.Todo) (models.Todo, error) {
	//service logic to create todo
	if todo.Title == "" {
		return models.Todo{}, errors.New("Title is required !!")
	}
	//also add a logic of adding automatically the Done is false

	err := repository.CreateTodo(todo)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil

}

func GetAllTodos() ([]primitive.M, error) {
	//service logic to get all todos
	todos, err := repository.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func DeleteTodo(id string) error {
	return repository.DeleteTodo(id)
}

func UpdateTodo(id string , updatedData models.Todo) error{
	return repository.UpdateTodo(id,updatedData)
}