package service

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/Ndraaa15/cordova/api/authentication/repository"
	"github.com/Ndraaa15/cordova/config/email"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/errors"
)

type AuthServiceImpl interface {
	ValidateAccount(c context.Context, token string, client *auth.Client) (*domain.User, error)
	RegisterAccount(c context.Context, req *domain.SignupRequest, client *auth.Client) (string, error)
}

type AuthService struct {
	ar repository.AuthRepositoryImpl
}

func NewAuthService(authRepository *repository.AuthRepositoryImpl) AuthServiceImpl {
	return &AuthService{authRepository}
}

func (as *AuthService) ValidateAccount(c context.Context, token string, client *auth.Client) (*domain.User, error) {
	//check existing id in database & check isActive
	resp, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}
	user, err := client.GetUser(c, resp.UID)
	if err != nil {
		return nil, err
	}

	if !user.EmailVerified {
		return nil, errors.ErrUserNotVerified
	}

	fmt.Println(*resp)
	//if there no exist create new one
	return &domain.User{}, nil
}

func (as *AuthService) RegisterAccount(c context.Context, req *domain.SignupRequest, authClient *auth.Client) (string, error) {
	//check existing email in database

	if req.Password != req.ConfirmPassword {
		return "", errors.ErrPasswordNotSame
	}

	resp, err := authClient.CreateUser(c, (&auth.UserToCreate{}).DisplayName(req.Name).Email(req.Email).Password(req.Password))

	if err != nil {
		return "", errors.ErrFailedCreateAccount
	}

	fmt.Println(resp.UID)

	//adding to database

	link, err := authClient.EmailVerificationLink(context.Background(), req.Email)
	if err != nil {
		return "", errors.ErrFailedCreateAccount
	}

	go sendEmailVerification(req, link)

	return resp.UID, nil
}

func sendEmailVerification(req *domain.SignupRequest, link string) error {
	mailClient := email.NewEmailClient()
	mailClient.SetSender("fuwafu212@gmail.com")
	mailClient.SetReciever(req.Email)
	mailClient.SetSubject("Email Verification")
	if err := mailClient.SetBodyHTML(req.Name, link); err != nil {
		return err
	}
	if err := mailClient.SendMail(); err != nil {
		return err
	}
	return nil
}
