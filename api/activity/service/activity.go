package service

import "github.com/Ndraaa15/cordova/api/activity/repository"

type ActivityServiceImpl interface {
}

type ActivityService struct {
	cr repository.ActivityRepositoryImpl
}

func NewActivityService(activityRepository *repository.ActivityRepositoryImpl) ActivityServiceImpl {
	return &ActivityService{activityRepository}
}
