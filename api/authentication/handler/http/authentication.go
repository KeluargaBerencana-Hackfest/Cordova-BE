package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/Ndraaa15/cordova/api/authentication/service"
	"github.com/Ndraaa15/cordova/config/firebase"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/middleware"
	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/Ndraaa15/cordova/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	as         service.AuthServiceImpl
	authClient *auth.Client
}

func NewAuthHandler(authService service.AuthServiceImpl) *AuthHandler {
	ah := &AuthHandler{
		as: authService,
	}

	app, err := firebase.InitFirebase()
	if err != nil {
		log.Printf("[cordova-authentication] failed to initialize firebase client. Error : %v\n", err)
		return nil
	}

	ah.authClient = app.AuthFirebase()
	return ah
}

func (ah *AuthHandler) Mount(s *gin.RouterGroup) {
	auth := s.Group("/auth")
	auth.POST("/signin", middleware.ValidateJWTToken(), ah.Login)
	auth.POST("/signup", ah.Register)
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
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
			log.Printf("[cordova-authentication] failed to sign in  user. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-authentication] sucess to sign in user.")
		response.Success(ctx, code, message, data)
	}()

	id, exist := ctx.Get("user")
	if !exist {
		return
	}

	res, err := ah.as.ValidateAccount(c, id.(string), ah.authClient)
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
		message = "Success to login"
		data = res
	}
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
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
			log.Printf("[cordova-authentication] failed to register new user. Error : %v\n", err)
			response.Error(ctx, code, err, message, nil)
			return
		}
		log.Printf("[cordova-authentication] success to register new user. ID : %v\n", data)
		response.Success(ctx, code, message, data)
	}()

	req := &domain.SignupRequest{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		code = http.StatusBadRequest
	}

	res, err := ah.as.RegisterAccount(c, req, ah.authClient)
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
		data = res
	}
}
