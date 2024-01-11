package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	AuthHandler "github.com/Ndraaa15/cordova/api/authentication/handler/http"
	AuthRepository "github.com/Ndraaa15/cordova/api/authentication/repository"
	AuthService "github.com/Ndraaa15/cordova/api/authentication/service"
	UserHandler "github.com/Ndraaa15/cordova/api/user/handler/http"
	UserRepository "github.com/Ndraaa15/cordova/api/user/repository"
	UserService "github.com/Ndraaa15/cordova/api/user/service"
	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/middleware"
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type Server struct {
	engine   *gin.Engine
	server   *http.Server
	handlers []Handler
}

type Handler interface {
	Mount(engine *gin.RouterGroup)
}

func NewServer() (*Server, error) {
	s := &Server{
		engine: gin.Default(),
	}

	s.server = &http.Server{
		Handler:      s.engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	db, err := database.ConnDatabase()
	if err != nil {
		log.Printf("[cordova-server] failed to initialize connection to postgres database. Error : %v\n", err)
		return nil, err
	}

	if err := db.MigrateDatabase(); err != nil {
		log.Printf("[cordova-server] failed to migrate schema to postgres database. Error : %v\n", err)
		return nil, err
	}

	authRepository := AuthRepository.NewAuthRepository(db)
	authService := AuthService.NewAuthService(authRepository)
	authHandler := AuthHandler.NewAuthHandler(authService)

	userRepository := UserRepository.NewUserRepository(db)
	userService := UserService.NewUserService(userRepository)
	userHandler := UserHandler.NewUserHandler(userService)

	s.handlers = []Handler{authHandler, userHandler}

	return s, nil
}

func (s *Server) StartServer() {
	s.engine.Use(middleware.CORS())
	for _, handler := range s.handlers {
		handler.Mount(s.engine.Group("/api/v1"))
	}
}

func RunServer() int {
	s, err := NewServer()

	if err != nil {
		return ErrBadConfig
	}

	s.StartServer()

	if err := s.engine.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
		return ErrInternalServer
	}

	return CodeSuccess
}
