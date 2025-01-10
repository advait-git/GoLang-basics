package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice class")

	// var fruitlist = []string{"Apple","Peach"}
	// fmt.Printf("%T \n",fruitlist)
	
	// fruitlist = append(fruitlist,"Mango","Banana")
	// fmt.Println("Fruit List are",fruitlist)

	// fruitlist = append(fruitlist[1:3])  
	// //the upper bracket no. is not inclusive
	// fmt.Println(fruitlist)

	highScore := make([]int,4)
	highScore[0]=234
	highScore[1]=999
	highScore[2]=250
	highScore[3]=300

	highScore=append(highScore,555,666,777)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)


	//how to remove a value from slice

	var courses=[]string{"reactjs","js","java","ruby","python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}