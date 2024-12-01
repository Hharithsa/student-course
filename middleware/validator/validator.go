package validator

import (
	"net/http"

	"github.com/Hharithsa/student-course-registration/entity"
	"github.com/gin-gonic/gin"
)

func ValidateStudentRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student entity.Student

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ValidateCoursesRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course entity.Course

		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
