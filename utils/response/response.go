package response

import "github.com/gin-gonic/gin"

type response struct {
	Status  status      `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"isSuccess"`
}

func Success(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, response{
		Status: status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Error:   "-",
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, err error, message string, data interface{}) {
	ctx.JSON(code, response{
		Status: status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Error:   err.Error(),
		Data:    data,
	})
}