package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	advait := User{"Dpaa", "advaittiwari2@gmail.com", true, 22}
	fmt.Println(advait)
	fmt.Printf("%+v\n", advait)
	fmt.Printf("name is %v and email is %v\n", advait.Name, advait.Email)
	advait.GetStatus()
	advait.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}

func (u User) GetStatus(){
	fmt.Println("Status of the user is : ",u.Status)
}

func (u User) NewMail(){
	u.Email="test@go.dev"
	fmt.Println("Email of user is",u.Email)
}