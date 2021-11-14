package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Student struct (model)
type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
	Address *Address `json: "address"`
}

// Address struct (model)
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

var students []Student

//get all student list
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

//get student by id
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	for _, item := range students {
		if strconv.FormatInt(int64(item.Id), 10) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf(w, "No record found with this ID")
}

//create new student
func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	for _, item := range students {
		if item.Id == student.Id {
			fmt.Fprintf(w, "ID already Exist")
			return
		}
	}
	students = append(students, student)
	json.NewEncoder(w).Encode(students)
}

//update student by id
func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	for index, item := range students {
		if strconv.FormatInt(int64(item.Id), 10) == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			for _, item := range students {
				if item.Id == student.Id {
					fmt.Fprintf(w, "ID already Exist")
					return
				}
			}
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	fmt.Fprintf(w, "ID Not Found")
}

//delete student by id
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	for index, item := range students {
		if strconv.FormatInt(int64(item.Id), 10) == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to restful api")
	println("On HomePage")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api", homePage)
	myRouter.HandleFunc("/api/students", getStudents).Methods("GET")
	myRouter.HandleFunc("/api/students/{id}", getStudent).Methods("GET")
	myRouter.HandleFunc("/api/students", createStudent).Methods("POST")
	myRouter.HandleFunc("/api/students/{id}", updateStudent).Methods("PUT")
	myRouter.HandleFunc("/api/students/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

	students = append(students, Student{Id: 1, Name: "Riad Safowan", Age: 21, Courses: []string{"Bangla", "English"}, Address: &Address{City: "Dhaka", Country: "Bangladesh"}})
	students = append(students, Student{Id: 2, Name: "Riad Safowan", Age: 21, Courses: []string{"Bangla", "English"}, Address: &Address{City: "Dhaka", Country: "Bangladesh"}})
	students = append(students, Student{Id: 3, Name: "Riad Safowan", Age: 21, Courses: []string{"Bangla", "English"}, Address: &Address{City: "Dhaka", Country: "Bangladesh"}})
	students = append(students, Student{Id: 4, Name: "Riad Safowan", Age: 21, Courses: []string{"Bangla", "English"}, Address: &Address{City: "Dhaka", Country: "Bangladesh"}})

	handleRequest()

}
