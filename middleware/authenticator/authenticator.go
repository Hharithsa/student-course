package authenticator

import (
	"github.com/Hharithsa/student-course-registration/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey != config.Envs.APIKey {
			c.JSON(401, gin.H{"error": "Unauthorized API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
