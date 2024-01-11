package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/Ndraaa15/cordova/api/user/service"
	"github.com/Ndraaa15/cordova/config/firebase"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/Ndraaa15/cordova/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us         service.UserServiceImpl
	authClient *auth.Client
}

func NewUserHandler(userService service.UserServiceImpl) *UserHandler {
	uh := &UserHandler{
		us: userService,
	}

	app, err := firebase.InitFirebase()
	if err != nil {
		log.Printf("[cordova-user] failed to initialize firebase client. Error : %v\n", err)
		log.Fatal("failed to initialize firebase client")
	}

	uh.authClient = app.AuthFirebase()

	return uh
}

func (uh *UserHandler) Mount(s *gin.RouterGroup) {
	user := s.Group("/user")
	user.PUT("/update", uh.UpdateUser)
	user.POST("/update/photo-profile", uh.UploadPhotoProfile)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
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
			log.Printf("[cordova-user] failed to update user. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-user] success to update user.")
		response.Success(ctx, code, message, data)
	}()

	id, exist := ctx.Get("user")
	if !exist {
		return
	}

	req := &domain.UserUpdateRequest{}
	if err = ctx.ShouldBindJSON(req); err != nil {
		code = http.StatusBadRequest
		message = "Failed to register account"
		return
	}

	res, err := uh.us.UpdateUserData(c, req, id.(string), uh.authClient)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success update user"
		data = res
	}
}

func (uh *UserHandler) UploadPhotoProfile(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			log.Printf("[cordova-user] failed to update user photo profile. Error : %v\n", err)
			response.Error(ctx, code, err, message, data)
			return
		}
		log.Printf("[cordova-user] success to update user photo profile.")
		response.Success(ctx, code, message, data)
	}()

	// id, exist := ctx.Get("user")
	// if !exist {
	// 	return
	// }

	photoProfile, _, err := ctx.Request.FormFile("photo-profile")
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := uh.us.UploadPhoto(c, photoProfile, "7jZRFXJM4MQ1MDGIjNEUiqUBoBr1")
	if err != nil {
		code = http.StatusInternalServerError
		message = ""
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
	default:
		message = "Success to upload photo profile"
		data = res
	}
}
