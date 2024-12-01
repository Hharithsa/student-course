package api

import (
	"database/sql"

	courseHandler "github.com/Hharithsa/student-course-registration/handler/course"
	registrationHandler "github.com/Hharithsa/student-course-registration/handler/registration"
	studentHandler "github.com/Hharithsa/student-course-registration/handler/student"
	"github.com/Hharithsa/student-course-registration/middleware/authenticator"
	"github.com/Hharithsa/student-course-registration/middleware/logger"
	courseRepo "github.com/Hharithsa/student-course-registration/repository/course"
	registrationRepo "github.com/Hharithsa/student-course-registration/repository/registration"
	studentRepo "github.com/Hharithsa/student-course-registration/repository/student"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) Server {
	return &server{
		addr: addr,
		db:   db,
	}
}

func (s *server) Run() error {
	router := gin.New()

	s.addAllMiddleware(router)
	s.addAllRoutes(router)

	return router.Run(s.addr)
}

func (s *server) addAllRoutes(router *gin.Engine) {
	studentRepo := studentRepo.NewRepository(s.db)
	studentHandler := studentHandler.NewHandler(studentRepo)
	studentHandler.AddStudentsRoutes(router)

	courseRepo := courseRepo.NewRepository(s.db)
	courseHandler := courseHandler.NewHandler(courseRepo)
	courseHandler.AddCoursesRoutes(router)

	registrationRepo := registrationRepo.NewRepository(s.db)
	registrationHandler := registrationHandler.NewHandler(registrationRepo)
	registrationHandler.AddRegistrationsRoutes(router)
}

func (s *server) addAllMiddleware(router *gin.Engine) {
	router.Use(logger.LoggingMiddleware())
	router.Use(authenticator.AuthMiddleware())
}
