package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file class")

	content:="This is a file demo"

	file , err := os.Create("./mydemo.txt")

	if err!=nil{
		panic(err)
	}

	length , err := io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("length is : ",length)
	defer file.Close()
	readFile("./mydemo.txt")
}

func readFile(filename string){
	databyte , err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("text data inside the file:\n",string(databyte))
}

func checkNilError(err error){
	if err !=nil{
		panic(err)
	}
}