package student

import (
	"net/http"
	"strconv"

	"github.com/Hharithsa/student-course-registration/entity"
	"github.com/Hharithsa/student-course-registration/middleware/validator"
	"github.com/Hharithsa/student-course-registration/repository/student"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddStudentsRoutes(rg *gin.Engine)
}

type handler struct {
	repo student.Repository
}

func NewHandler(studentRepo student.Repository) Handler {
	return &handler{
		repo: studentRepo,
	}
}

func (h *handler) AddStudentsRoutes(rg *gin.Engine) {
	students := rg.Group("/students")

	students.POST("/", validator.ValidateStudentRequest(), h.postStudents)
	students.GET("/", h.getStudents)
	students.GET("/:id", h.getStudentsByID)
	students.PUT("/:id", validator.ValidateStudentRequest(), h.putStudentsByID)
	students.DELETE("/:id", h.deleteStudentsByID)
}

func (h *handler) postStudents(c *gin.Context) {
	body, exists := c.Get("student")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Student data not found in context",
		})
		return
	}

	student, ok := body.(entity.Student)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast student data",
		})
		return
	}

	if err := h.repo.CreateStudents(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student created successfully"})
}

func (h *handler) getStudents(c *gin.Context) {
	students, err := h.repo.QueryStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *handler) getStudentsByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := h.repo.QueryStudentsByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if student == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Student by such ID"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *handler) putStudentsByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, _ := h.repo.QueryStudentsByID(idInt)
	if s == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such Student ID"})
		return
	}

	body, exists := c.Get("student")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Student data not found in context",
		})
		return
	}

	student, ok := body.(entity.Student)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast student data",
		})
		return
	}
	student.ID = idInt

	if err := h.repo.UpdateStudents(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func (h *handler) deleteStudentsByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, _ := h.repo.QueryStudentsByID(idInt)
	if student == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such Student ID"})
		return
	}

	if err := h.repo.DeleteStudents(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
