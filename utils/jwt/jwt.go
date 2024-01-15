package jwt

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

func DecodeToken(client *auth.Client, token string) (*auth.Token, error) {
	resp, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
