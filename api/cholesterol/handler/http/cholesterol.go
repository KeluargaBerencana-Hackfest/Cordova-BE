package http

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/cordova/api/cholesterol/service"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/middleware"
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

func (ch *CholesterolHandler) Mount(s *gin.RouterGroup) {
	cholesterol := s.Group("/cholesterol")
	cholesterol.POST("/check", middleware.ValidateJWTToken(), ch.CheckCholesterol)
	cholesterol.GET("/history", middleware.ValidateJWTToken(), ch.GetCholesterolHistory)
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

	_, exist := ctx.Get("user")
	if !exist {
		code = http.StatusBadRequest
		message = "Failed to get user id"
		err = errors.ErrBadRequest
		return
	}

	req := &domain.CholesterolRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		code = http.StatusBadRequest
	}

	res, err := ch.cs.CalculateCholesterol(c, req)

	if err != nil {
		code = http.StatusBadRequest
		message = "Failed to calculate cholesterol"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success calculate cholesterol"
		data = res
	}
}

func (ch *CholesterolHandler) GetCholesterolHistory(ctx *gin.Context) {
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

	id, exist := ctx.Get("user")
	if !exist {
		code = http.StatusBadRequest
		message = "Failed to get user id"
		err = errors.ErrBadRequest
		return
	}

	res, err := ch.cs.GetCholesterolHistory(c, id.(string))
	if err != nil {
		code = http.StatusBadRequest
		message = "Failed to get cholesterol history"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success get cholesterol history"
		data = res
	}
}
