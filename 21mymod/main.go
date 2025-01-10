package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func main() {
	fmt.Println("hello world")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000",r))
}

func greeter(){
	fmt.Println("hey ther mod users")
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to golang class via server</h1>"))
}
