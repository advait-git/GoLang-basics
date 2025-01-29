package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/advait-git/todolist/api/v1"
	"github.com/advait-git/todolist/conf"
)

func main() {
	fmt.Println("Hello World")
	conf.ConnectDB()
	//called the router for calling the apis
	r := api.Router()


	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
