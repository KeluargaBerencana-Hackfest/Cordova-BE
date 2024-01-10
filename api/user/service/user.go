package service

import (
	"context"

	"github.com/Ndraaa15/cordova/api/user/repository"
	"github.com/Ndraaa15/cordova/domain"
)

type UserServiceImpl interface {
	UpdateUserData(c context.Context, req *domain.User, userId string) (*domain.User, error)
}

type UserService struct {
	ur repository.UserRepositoryImpl
}

func NewUserService(userRepository *repository.UserRepositoryImpl) UserServiceImpl {
	return &UserService{userRepository}
}

func (us *UserService) UpdateUserData(c context.Context, req *domain.User, userId string) (*domain.User, error) {
	return &domain.User{}, nil
}

func ValidateUpdateRequest(req *domain.UserUpdate, user *domain.User) *domain.User {
	if req.Name != "" {
		user.Name = req.Name
	}

	//another

	return user
}
