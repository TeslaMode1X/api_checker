package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "main/models"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var courses []m.Course

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].CourseId < courses[j].CourseId
	})

	json.NewEncoder(w).Encode(courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid course id", http.StatusBadRequest)
		return
	}

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	http.Error(w, "No course found with the given id", http.StatusNotFound)
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		http.Error(w, "Please send some data", http.StatusBadRequest)
		return
	}

	var course m.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		http.Error(w, "No data in JSON that you sent", http.StatusBadRequest)
		return
	}

	for _, v := range courses {
		if v.CourseName == course.CourseName {
			http.Error(w, "Course with this name already exists", http.StatusConflict)
			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = rand.Intn(1000) + 1 // Generate IDs greater than 0

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid course id", http.StatusBadRequest)
		return
	}

	for index, course := range courses {
		if course.CourseId == id {
			var updatedCourse m.Course
			if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
				http.Error(w, "Invalid input data", http.StatusBadRequest)
				return
			}
			if updatedCourse.IsEmpty() {
				http.Error(w, "No data in JSON that you sent", http.StatusBadRequest)
				return
			}
			updatedCourse.CourseId = id
			courses[index] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	http.Error(w, "Course not found", http.StatusNotFound)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid course id", http.StatusBadRequest)
		return
	}

	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	http.Error(w, "Course not found", http.StatusNotFound)
}
