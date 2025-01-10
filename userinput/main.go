package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welecome to the train"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")

	input , _ :=reader.ReadString('\n')
	fmt.Println("Thx for the jorney",input)
	fmt.Printf("Type of the rating is %T",input)
}
