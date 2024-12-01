package registration

import (
	"net/http"
	"strconv"

	"github.com/Hharithsa/student-course-registration/entity"
	"github.com/Hharithsa/student-course-registration/repository/registration"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddRegistrationsRoutes(rg *gin.Engine)
}

type handler struct {
	repo registration.Repository
}

func NewHandler(studentRepo registration.Repository) Handler {
	return &handler{
		repo: studentRepo,
	}
}

func (h *handler) AddRegistrationsRoutes(rg *gin.Engine) {
	registrations := rg.Group("/registrations")

	registrations.POST("/", h.postRegistrations)
	registrations.GET("/", h.getRegistrations)
	registrations.GET("/:student_id", h.getRegistrationsByStudentID)
}

func (h *handler) postRegistrations(c *gin.Context) {
	registration := entity.Registration{}

	if err := c.ShouldBindBodyWithJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateRegistration(registration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration created successfully"})
}

func (h *handler) getRegistrations(c *gin.Context) {
	registration, err := h.repo.QueryRegistrations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registration)
}

func (h *handler) getRegistrationsByStudentID(c *gin.Context) {
	id := c.Param("student_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registration, err := h.repo.QueryRegistrationsByStudentID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registration)
}
