package main

import "fmt"

func main() {
	fmt.Println("Class on pointers !!")

	// var ptr *int
	// fmt.Println("Value",ptr)

	myNumber := 23
	var ptr = &myNumber //reference means => &

	fmt.Println("Address",ptr)
	fmt.Println("Value",*ptr)
}