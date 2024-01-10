package http

import (
	"github.com/Ndraaa15/cordova/api/user/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	cs service.UserServiceImpl
}

func NewUserHandler(userService service.UserServiceImpl) *UserHandler {
	return &UserHandler{userService}
}

func (ah *UserHandler) Mount(s *gin.Engine) {
	user := s.Group("/user")
	user.PUT("/update/:id")
}

func (uh *UserHandler) UpdateUserData(ctx *gin.Context) {

}
