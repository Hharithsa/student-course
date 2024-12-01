package entity

import "time"

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required,gt=0,lte=120"`
	College string `json:"college" binding:"required"`
	Year    int    `json:"year" binding:"required,gte=1,lte=5"`
}

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required,min=1,max=120"`
	Description string `json:"description" binding:"required"`
}

type Registration struct {
	StudentID int       `json:"studentId" binding:"required"`
	CourseID  int       `json:"courseId" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
