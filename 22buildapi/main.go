package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for courses -file

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware , helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to build api class")
	r := mux.NewRouter()

	//seeding of data
	courses = append(courses, Course{CourseId: "2" , CourseName:"ReactJS",CoursePrice: 299,Author: &Author{Fullname: "Advait Tiwari",Website: "github.com/advait-git"}})
	courses = append(courses, Course{CourseId: "3" , CourseName:"GoLang",CoursePrice: 499,Author: &Author{Fullname: "Akshita",Website: "github.com/dpaa-git"}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":8000",r))
}

//controllers -file

//serve home route or health check

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the GoLang Class</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from the req
	params := mux.Vars(r)

	//loop through the course, find the matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found of this id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//wht if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please put some value inside the json")
		return
	}

	//Todo : check i the course name is duplicate
	//loop , course_name,json response that it is duplicate
	for _ , exsistingcourse :=range courses{
		if exsistingcourse.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("ERROR : You are using duplicate values inside the json")
			return
		}
	}

	//generate a unique id , and also convert them to string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	//first grab id from req
	params := mux.Vars(r)

	//loop throught the value , once we get id , remove , add value with my ID which we will get from param
	for index , course :=range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")

	param :=mux.Vars(r)

	//loop , id , remove (index,index+1)
	for index,course := range courses{
		if course.CourseId == param["id"]{
			courses = append(courses[:index],courses[index+1:]... )
			json.NewEncoder(w).Encode("You have successfully deleted the course !")
			break
		}
	}
}