package http

import "github.com/gin-gonic/gin"

type GoogleHandler struct{}

func NewGoogleHandler() *GoogleHandler {
	return &GoogleHandler{}
}

func (gh *GoogleHandler) Mount(engine *gin.Engine) {
	engine.GET("/google", )
}
