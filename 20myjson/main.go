package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course_name"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welvome to myjson")
	DecodeJson()

}

func EncodeJson() {
	ytCourses := []course{
		{"React Js", 100, "YT.com", "abc", []string{"yt", "frontend"}},
		{"Java Springboot", 100, "YT/dpaa.com", "abc123", []string{"java", "backend"}},
		{"GO lang", 500, "YT.com", "abc1234", nil},
	}

	finalJson, err := json.MarshalIndent(ytCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}


func DecodeJson(){
	jsonDataFromWeb :=[]byte(`
		{
                "course_name": "Java Springboot",
                "Price": 100,
                "website": "YT/dpaa.com",
                "tags": [
                        "java",
                        "backend"
                ]
        }
	`)

	var lcoCourse course

	checkVaild := json.Valid(jsonDataFromWeb)

	if checkVaild{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("that json was not valid")
	}

	//sometimes you just want to add data by key value pair and not use struct

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k,v :=range myOnlineData{
		fmt.Printf("Key is %v and value is %v\n",k,v)
	}
}
