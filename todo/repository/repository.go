package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/advait-git/todolist/conf"
	"github.com/advait-git/todolist/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(todo models.Todo) error {
	//repo logic to create todo

	//create a context and made the connection with the mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := conf.GetCollection("tasks")

	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetAllTodos() ([]primitive.M, error) {
	//repo logic to get all todos

	//create a context and made the connection with the mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := conf.GetCollection("tasks")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var todos []primitive.M

	for cursor.Next(ctx) {
		var todo bson.M
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)

	}
	return todos, nil
}

func DeleteTodo(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := conf.GetCollection("tasks")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func UpdateTodo(id string, updatedData models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := conf.GetCollection("tasks")

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": updatedData}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated : ", updatedData)
	return nil
}
