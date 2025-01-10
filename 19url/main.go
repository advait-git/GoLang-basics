package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asd23asd3"

func main() {
	fmt.Println("Welcome to class of url")
	fmt.Println("myurl : ",myurl)

	//parsing 
	result , _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)

	qparam := result.Query()
	fmt.Printf("%T\n",qparam)
	fmt.Println(qparam["coursename"])

	for _ , val := range qparam{
		fmt.Println("the params : ",val)
	}

	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "local.dev",
		Path: "/dev",
		RawPath: "user=advait",
	}
	fmt.Println(partsofUrl.String())
}