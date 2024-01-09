package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	googleHandler "github.com/Ndraaa15/cordova/api/oauth/handler/http"
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type Server struct {
	engine *gin.Engine
	server *http.Server
	handlers []Handler
}

type Handler interface {
	Mount(engine *gin.Engine)
}

func NewServer() (*Server, error) {
	s := &Server{
		engine: gin.Default(),
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	s.handlers = []Handler{googleHandler.NewGoogleHandler()}


	return s, nil
}

func (s *Server) StartServer(){
	for _, handler := range s.handlers {
		handler.Mount(s.engine)
	}
} 

func RunServer() int {
	s, err := NewServer()

	if err != nil {
		return ErrBadConfig
	}

	s.StartServer()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT_ADDR")), s.engine); err != nil && err != http.ErrServerClosed {
		return ErrInternalServer
	}

	return CodeSuccess
}