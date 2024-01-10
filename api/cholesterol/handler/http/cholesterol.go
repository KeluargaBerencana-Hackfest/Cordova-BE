package http

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/cordova/api/cholesterol/service"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/Ndraaa15/cordova/utils/response"
	"github.com/gin-gonic/gin"
)

type CholesterolHandler struct {
	cs service.CholesterolServiceImpl
}

func NewCholesterolHandler(cholesterolService service.CholesterolServiceImpl) *CholesterolHandler {
	return &CholesterolHandler{cholesterolService}
}

func (ch *CholesterolHandler) Mount(s *gin.Engine) {
	_ = s.Group("/cholesterol/:id")
}

func (ch *CholesterolHandler) CheckCholesterol(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var (
		message string
		err     error
		data    interface{}
		code    = http.StatusOK
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, nil)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	req := &domain.CholesterolRequest{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		code = http.StatusBadRequest
	}

	if err != nil {
		code = http.StatusBadRequest
		message = "Failed to register account"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Please verified the email"
	}
}
