package main

import "fmt"

func main() {
	fmt.Println("Welcome to map class")

	language := make(map[string]string)
	// map[key]value

	language["JS"]="javascript"
	language["Py"]="python"
	language["rb"]="ruby"

	fmt.Println(language)
	fmt.Println(language["Py"])

	delete(language,"rb")
	fmt.Println("new language list",language)

	//loops in maps

	for key , value := range language{
		fmt.Printf("For key %v and for value %v \n",key,value)
	}

}