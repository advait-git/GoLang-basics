package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops class")
	// days:= []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}
	// fmt.Println(days)
//type 1
	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }
//type 2
	// for i:=range days{
	// 	fmt.Println(days[i])
	// }
//type 3 : for each loop
	// for idx,day :=range days{
	// 	fmt.Printf("index %v and days %v\n",idx,day)
	// }

	rougeValue:= 1
	for rougeValue < 10{
		if rougeValue==5{
			break
		}
		fmt.Println("value : ",rougeValue)
		rougeValue++
	}

}