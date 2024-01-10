package http

import (
	"github.com/Ndraaa15/cordova/api/cholesterol/service"
	"github.com/gin-gonic/gin"
)

type CholesterolHandler struct {
	cs service.CholesterolServiceImpl
}

func NewCholesterolHandler(cholesterolService service.CholesterolServiceImpl) *CholesterolHandler {
	return &CholesterolHandler{cholesterolService}
}

func (ch *CholesterolHandler) Mount(s *gin.Engine) {
	_ = s.Group("/cholesterol")

}
