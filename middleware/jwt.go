package middleware

import (
	"net/http"
	"strings"

	"github.com/Ndraaa15/cordova/config/firebase"
	"github.com/Ndraaa15/cordova/utils/jwt"
	"github.com/gin-gonic/gin"
)

func ValidateJWTToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		client, err := firebase.InitFirebase()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error initialize firebase"})
			return
		}

		header := ctx.Request.Header.Get("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You must be logged in first."})
			return
		}

		tokenParts := strings.SplitN(header, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		tokenString := tokenParts[1]

		claims, err := jwt.DecodeToken(client.AuthFirebase(), tokenString)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		userID := claims.UID

		ctx.Set("user", userID)
		ctx.Next()
	}
}
