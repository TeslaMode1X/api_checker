package models

type Course struct {
	CourseId    int     `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
