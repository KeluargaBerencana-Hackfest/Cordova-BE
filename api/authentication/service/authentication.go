package service

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/Ndraaa15/cordova/api/authentication/repository"
	"github.com/Ndraaa15/cordova/config/email"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/errors"
)

type AuthServiceImpl interface {
	ValidateAccount(c context.Context, id string, client *auth.Client) (*domain.User, error)
	RegisterAccount(c context.Context, req *domain.SignupRequest, client *auth.Client) (string, error)
}

type AuthService struct {
	ar repository.AuthRepositoryImpl
}

func NewAuthService(authRepository repository.AuthRepositoryImpl) AuthServiceImpl {
	return &AuthService{authRepository}
}

func (as *AuthService) ValidateAccount(c context.Context, id string, authClient *auth.Client) (*domain.User, error) {
	user, err := authClient.GetUser(c, id)
	if err != nil {
		return nil, errors.ErrFailedGetAccount
	}

	if !user.EmailVerified {
		return nil, errors.ErrUserNotVerified
	}

	count, err := as.ar.CountEmailAccount(&c, user.Email)
	if err != nil {
		return nil, errors.ErrFailedCountEmailUser
	}

	if count != 0 {
		userRecord, err := as.ar.GetAccountByID(c, user.UID)
		if err != nil {
			return nil, errors.ErrFailedGetAccount
		}
		return userRecord, nil
	}

	newUser, err := as.ar.SaveAccount(c, parseUserReq(user))
	if err != nil {
		return nil, errors.ErrFailedSaveAccount
	}

	return newUser, nil
}

func (as *AuthService) RegisterAccount(c context.Context, req *domain.SignupRequest, authClient *auth.Client) (string, error) {
	count, err := as.ar.CountEmailAccount(&c, req.Email)
	if err != nil {
		return "", errors.ErrFailedCountEmailUser
	}

	if count != 0 {
		return "", errors.ErrEmailAlreadyExist
	}

	if req.Password != req.ConfirmPassword {
		return "", errors.ErrPasswordNotSame
	}

	resp, err := authClient.CreateUser(c, (&auth.UserToCreate{}).DisplayName(req.Name).Email(req.Email).Password(req.Password))
	if err != nil {
		return "", errors.ErrFailedCreateAccount
	}

	_, err = as.ar.SaveAccount(c, parseUserReq(resp))
	if err != nil {
		return "", err
	}

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

// Create user account for db if there no existing email (because the user can login via google)
func parseUserReq(oauth *auth.UserRecord) *domain.User {
	return &domain.User{
		ID:    oauth.UID,
		Name:  oauth.DisplayName,
		Email: oauth.Email,
	}
}
