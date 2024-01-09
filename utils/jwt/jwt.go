package jwt

import (
	"os"
	"time"

	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/golang-jwt/jwt/v4"
)

func EncodeToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
		"id":  user.ID,
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", errors.ErrSigningJWT
	}
	return signedToken, nil
}

func DecodeToken(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if !decoded.Valid {
		return nil, errors.ErrClaimsJWT
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.ErrClaimsJWT
	}

	return claims, nil
}