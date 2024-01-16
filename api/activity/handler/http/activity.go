package http

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/cordova/api/activity/service"
	"github.com/Ndraaa15/cordova/middleware"
	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/Ndraaa15/cordova/utils/response"
	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	cs service.ActivityServiceImpl
}

func NewActivityHandler(activityService service.ActivityServiceImpl) *ActivityHandler {
	return &ActivityHandler{activityService}
}

func (ah *ActivityHandler) Mount(s *gin.RouterGroup) {
	activity := s.Group("/activity")
	activity.PUT("/checklist/:sub_activity_id", middleware.ValidateJWTToken(), ah.ChecklistActivity)
	activity.PUT("/unchecklist/:sub_activity_id", middleware.ValidateJWTToken(), ah.UnchecklistActivity)
	activity.GET("/all", middleware.ValidateJWTToken(), ah.GetAllActivity)
	activity.PUT("/regenerate", middleware.ValidateJWTToken(), ah.RegenerateActivity)
}

func (ah *ActivityHandler) ChecklistActivity(ctx *gin.Context) {
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
			log.Printf("[cordova-activity-http] failed to checklist. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-activity-http] success to checklist")
		response.Success(ctx, code, message, data)
	}()

	subActivityIDParam := ctx.Param("sub_activity_id")
	subActivityID, err := strconv.Atoi(subActivityIDParam)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	res, err := ah.cs.ChecklistActivity(c, subActivityID)
	if err != nil {
		code = http.StatusInternalServerError
		message = "Failed to checklist"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success to checklist"
		data = res
	}
}

func (ah *ActivityHandler) UnchecklistActivity(ctx *gin.Context) {
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
			log.Printf("[cordova-activity-http] failed to unchecklist. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-activity-http] success to unchecklist")
		response.Success(ctx, code, message, data)
	}()

	subActivityIDParam := ctx.Param("sub_activity_id")
	subActivityID, err := strconv.Atoi(subActivityIDParam)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	res, err := ah.cs.UnchecklistActivity(c, subActivityID)
	if err != nil {
		code = http.StatusInternalServerError
		message = "Failed to unchecklist"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success to unchecklist"
		data = res
	}
}

func (ah *ActivityHandler) GetAllActivity(ctx *gin.Context) {
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
			log.Printf("[cordova-activity-http] failed to get all activity. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-activity-http] success to get all activity")
		response.Success(ctx, code, message, data)
	}()

	id, exist := ctx.Get("user")
	if !exist {
		code = http.StatusBadRequest
		message = "Failed to get user id"
		err = errors.ErrBadRequest
		return
	}

	res, err := ah.cs.GetAllActivity(c, id.(string))
	if err != nil {
		code = http.StatusInternalServerError
		message = "Failed to get all activity"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success to get all activity"
		data = res
	}
}

func (ah *ActivityHandler) RegenerateActivity(ctx *gin.Context) {
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
			log.Printf("[cordova-activity-http] failed to regenerate activity. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-activity-http] success to regenerate activity")
		response.Success(ctx, code, message, data)
	}()

	id, exist := ctx.Get("user")
	if !exist {
		code = http.StatusBadRequest
		message = "Failed to get user id"
		err = errors.ErrBadRequest
		return
	}

	res, err := ah.cs.RegenerateActivity(c, id.(string))
	if err != nil {
		code = http.StatusInternalServerError
		message = "Failed regenerate activity"
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success regenerate activity"
		data = res
	}
}
