package logger

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	f, _ := os.OpenFile("logging.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	multiWriter := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multiWriter)
	return func(c *gin.Context) {
		log.Printf("Request: %s %s %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())

		writer := &responseWriter{
			ResponseWriter: c.Writer,
			statusCode:     200,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		log.Printf("Response: %d %s %s %s",
			writer.statusCode,
			writer.body.String(),
			c.Request.Method,
			c.Request.URL.Path)
	}
}

type responseWriter struct {
	gin.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(p []byte) (n int, err error) {
	rw.body.Write(p)
	return rw.ResponseWriter.Write(p)
}
