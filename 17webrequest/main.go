package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url="https://chaicode.com"

func main() {
	fmt.Println("Welcome to class")
	response , err :=http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response type %T \n",response)
	defer response.Body.Close() // to close the connection

	databyte , err := ioutil.ReadAll(response.Body)

	if err !=nil{
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}