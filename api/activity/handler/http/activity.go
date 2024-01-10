package http

import (
	"github.com/Ndraaa15/cordova/api/activity/service"
	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	cs service.ActivityServiceImpl
}

func NewActivityHandler(activityService service.ActivityServiceImpl) *ActivityHandler {
	return &ActivityHandler{activityService}
}

func (ah *ActivityHandler) Mount(s *gin.Engine) {
	_ = s.Group("/activity")
}
