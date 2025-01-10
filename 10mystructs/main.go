package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	advait := User{"Dpaa","advaittiwari2@gmail.com",true,22}
	fmt.Println(advait)
	fmt.Printf("%+v\n",advait)
	fmt.Printf("name is %v and email is %v",advait.Name,advait.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
