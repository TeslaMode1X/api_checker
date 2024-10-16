package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	h "main/handlers"
	"net/http"
)

func main() {
	fmt.Println("API check")

	r := mux.NewRouter()

	r.HandleFunc("/", h.ServeHome).Methods("GET")
	r.HandleFunc("/courses", h.GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", h.GetCourse).Methods("GET")
	r.HandleFunc("/course", h.CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", h.UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", h.DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4040", r))
}
