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

// create struct for api

type Course struct {
	CourseName string  `json:"course_name"`
	CourseId   string  `json:"course_id"`
	Website    string  `json:"website"`
	Author     *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
}

var courses []Course

func main() {
	fmt.Println("Server running on port 8081")
	r := mux.NewRouter()
	r.HandleFunc("/", testApi).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{courseId}", getCourseById).Methods("GET")
	r.HandleFunc("/course/{courseId}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{courseId}", deleteCourseById).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// middleware func
func (body *Course) isId() bool {
	return body.CourseId == ""
}

func (body *Course) isBodyCheck() bool {
	return body.CourseName == ""
}

func testApi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Test Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(courses)
	return
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get params
	params := mux.Vars(r)
	var isCourseFind bool = false
	// find course by id
	for _, course := range courses {
		if course.CourseId == params["courseId"] {
			isCourseFind = true
			_ = json.NewEncoder(w).Encode(course)
			return
		}
	}
	if isCourseFind == false {
		_ = json.NewEncoder(w).Encode("The course not found")
		return
	}

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		fmt.Println(r.Body)
		json.NewEncoder(w).Encode("The body should not be an empty")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isBodyCheck() {
		json.NewEncoder(w).Encode("The course name is required")
		return
	}
	// create unique id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	_ = json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode("The body should not be an empty")
	}
	params := mux.Vars(r)
	var isUpdate bool = false
	//TODO: find course after remove and update
	for index, value := range courses {
		if value.CourseId == params["courseId"] {
			isUpdate = true
			remainingCourses := append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["courseId"]
			_ = append(remainingCourses, course)
			_ = json.NewEncoder(w).Encode(course)
			return

		}
	}

	if isUpdate == false {
		_ = json.NewEncoder(w).Encode("The course not a found")
	}

}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params == nil {
		_ = json.NewEncoder(w).Encode("The course not found")
	}
	var isDelete bool = false
	for index, value := range courses {
		if value.CourseId == params["courseId"] {
			isDelete = true
			remainingCourses := append(courses[:index], courses[index+1:]...)
			courses = remainingCourses
			_ = json.NewEncoder(w).Encode("The course is deleted")
			return
		}
	}
	if isDelete == false {
		_ = json.NewEncoder(w).Encode("Failed to delete the course")
		return
	}

}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	courses = append(courses[0:0])
	_ = json.NewEncoder(w).Encode("Courses are deleted")
	return
}
