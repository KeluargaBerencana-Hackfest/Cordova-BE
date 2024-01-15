package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/Ndraaa15/cordova/api/user/repository"
	"github.com/Ndraaa15/cordova/config/firebase"
	"github.com/Ndraaa15/cordova/domain"
)

type UserServiceImpl interface {
	UpdateUserData(c context.Context, req *domain.UserUpdateRequest, userId string, authClient *auth.Client) (*domain.User, error)
	UploadPhoto(ctx context.Context, file multipart.File, id string) (*domain.User, error)
}

type UserService struct {
	ur repository.UserRepositoryImpl
}

func NewUserService(userRepository repository.UserRepositoryImpl) UserServiceImpl {
	return &UserService{userRepository}
}

func (us *UserService) UpdateUserData(c context.Context, req *domain.UserUpdateRequest, userId string, authClient *auth.Client) (*domain.User, error) {
	user, err := us.ur.GetUserByID(c, userId)
	if err != nil {
		log.Printf("[cordova-user-service] failed to get user from database. Error : %v\n", err)
		return nil, err
	}

	//change password in front end

	if req.Email != "" && user.Email != req.Email {
		_, err := authClient.UpdateUser(c, userId, (&auth.UserToUpdate{}).Email(req.Email))
		if err != nil {
			log.Printf("[cordova-user-service] failed to update user (email) from firebase. Error : %v\n", err)
			return nil, err
		}
	}

	if req.Name != "" && user.Name != req.Name {
		_, err := authClient.UpdateUser(c, userId, (&auth.UserToUpdate{}).DisplayName(req.Name))
		if err != nil {
			log.Printf("[cordova-user-service] failed to update user (name) from firebase. Error : %v\n", err)
			return nil, err
		}
	}

	user, err = validateUpdateRequest(req, user)
	if err != nil {
		log.Printf("[cordova-user-service] failed to validate update request. Error : %v\n", err)
		return nil, err
	}

	user, err = us.ur.UpdateUser(c, user)
	if err != nil {
		log.Printf("[cordova-user-service] failed to update user from database. Error : %v\n", err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) UploadPhoto(ctx context.Context, file multipart.File, id string) (*domain.User, error) {
	user, err := us.ur.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("[cordova-user-service] failed to get user from database. Error : %v\n", err)
		return nil, err
	}

	formByte, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[cordova-user-service] failed to read file. Error : %v\n", err)
		return nil, err
	}

	photo, err := uploadPhotos(ctx, formByte, fmt.Sprintf("photo-%s", user.ID))
	if err != nil {
		log.Printf("[cordova-user-service] failed to upload photo profile. Error : %v\n", err)
		return nil, err
	}

	user.PhotoProfile = photo

	user, err = us.ur.UpdateUser(ctx, user)
	if err != nil {
		log.Printf("[cordova-user-service] failed to update from database. Error : %v\n", err)
		return nil, err
	}

	return user, nil
}

func validateUpdateRequest(req *domain.UserUpdateRequest, user *domain.User) (*domain.User, error) {
	if req.Name != "" && req.Name != user.Name {
		user.Name = req.Name
	}

	if req.Email != "" && req.Email != user.Email {
		user.Email = req.Email
	}

	if req.Birthday != "" && req.Birthday != user.Birthday.Format("2006-01-02") {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			user.Birthday = parsedBirthday
		} else {
			return nil, err
		}
	}

	if req.Gender != user.Gender {
		user.Gender = req.Gender
	}

	if req.Weight != user.Weight && req.Weight > 0 {
		user.Weight = req.Weight
	}

	if req.Height != user.Height && req.Height > 0 {
		user.Height = req.Height
	}

	if req.Exercise != user.Exercise {
		user.Exercise = req.Exercise
	}

	if req.PhysicalActivity != user.PhysicalActivity && req.PhysicalActivity > 0 {
		user.PhysicalActivity = req.PhysicalActivity
	}

	if req.SleepHours != user.SleepHours && req.SleepHours > 0 {
		user.SleepHours = req.SleepHours
	}

	if req.Smoking != user.Smoking {
		user.Smoking = req.Smoking
	}

	if req.AlcoholConsumption != user.AlcoholConsumption {
		user.AlcoholConsumption = req.AlcoholConsumption
	}

	if req.SedentaryHours != user.SedentaryHours && req.SedentaryHours > 0 {
		user.SedentaryHours = req.SedentaryHours
	}

	if req.Diabetes != user.Diabetes {
		user.Diabetes = req.Diabetes
	}

	if req.FamilyHistory != user.FamilyHistory {
		user.FamilyHistory = req.FamilyHistory
	}

	if req.PreviousHeartProblem != user.PreviousHeartProblem {
		user.PreviousHeartProblem = req.PreviousHeartProblem
	}

	if req.MedicationUse != user.MedicationUse {
		user.MedicationUse = req.MedicationUse
	}

	if req.StressLevel != user.StressLevel && req.StressLevel > 0 {
		user.StressLevel = req.StressLevel
	}

	return user, nil
}

func uploadPhotos(ctx context.Context, file []byte, fileName string) (string, error) {
	f, err := firebase.InitFirebase()
	if err != nil {
		return "", err
	}
	link, err := f.UploadFile(ctx, file, fileName)
	if err != nil {
		return "", err
	}
	return link, nil
}
