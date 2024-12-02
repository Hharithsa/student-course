package course

import (
	"net/http"
	"strconv"

	"github.com/Hharithsa/student-course-registration/entity"
	"github.com/Hharithsa/student-course-registration/middleware/validator"
	"github.com/Hharithsa/student-course-registration/repository/course"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddCoursesRoutes(rg *gin.Engine)
}

type handler struct {
	repo course.Repository
}

func NewHandler(courseRepo course.Repository) Handler {
	return &handler{
		repo: courseRepo,
	}
}

func (h *handler) AddCoursesRoutes(rg *gin.Engine) {
	courses := rg.Group("/courses")

	courses.POST("/", validator.ValidateCoursesRequest(), h.postCourses)
	courses.GET("/", h.getCourses)
	courses.GET("/:id", h.getCoursesByID)
	courses.PUT("/:id", validator.ValidateCoursesRequest(), h.putCoursesByID)
	courses.DELETE("/:id", h.deleteCoursesByID)
}

func (h *handler) postCourses(c *gin.Context) {
	body, exists := c.Get("course")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Course data not found in context",
		})
		return
	}

	course, ok := body.(entity.Course)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast student data",
		})
		return
	}

	if err := h.repo.CreateCourses(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "course created successfully"})
}

func (h *handler) getCourses(c *gin.Context) {
	courses, err := h.repo.QueryCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *handler) getCoursesByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.repo.QueryCoursesByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if course == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Course by such ID"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *handler) putCoursesByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	co, _ := h.repo.QueryCoursesByID(idInt)
	if co == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such Course ID"})
		return
	}

	body, exists := c.Get("course")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Course data not found in context",
		})
		return
	}

	course, ok := body.(entity.Course)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast Course data",
		})
		return
	}
	course.ID = idInt

	if err := h.repo.UpdateCourses(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (h *handler) deleteCoursesByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, _ := h.repo.QueryCoursesByID(idInt)
	if course == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such course ID"})
		return
	}

	if err := h.repo.DeleteCourses(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
