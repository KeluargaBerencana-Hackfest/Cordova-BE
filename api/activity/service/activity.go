package service

import (
	"context"
	"log"
	"time"

	"github.com/Ndraaa15/cordova/api/activity/repository"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/activity"
	"github.com/Ndraaa15/cordova/utils/errors"
)

type ActivityServiceImpl interface {
	ChecklistActivity(c context.Context, activityID int) ([]*domain.ActivityDB, error)
	UnchecklistActivity(c context.Context, activityID int) ([]*domain.ActivityDB, error)
	GetAllActivity(c context.Context, userID string) ([]*domain.ActivityDB, error)
	RegenerateActivity(c context.Context, userID string) ([]*domain.ActivityDB, error)
}

type ActivityService struct {
	ar repository.ActivityRepositoryImpl
}

func NewActivityService(activityRepository repository.ActivityRepositoryImpl) ActivityServiceImpl {
	return &ActivityService{activityRepository}
}

func (as *ActivityService) GetAllActivity(c context.Context, userID string) ([]*domain.ActivityDB, error) {
	activities, err := as.ar.GetAllActivity(c, userID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get all activity. Error : %v\n", err)
		return nil, err
	}

	return activities, nil
}

func (as *ActivityService) RegenerateActivity(c context.Context, userID string) ([]*domain.ActivityDB, error) {
	userCholesterol, err := as.ar.GetUserCholesterolByID(c, userID, int(time.Now().Month()), time.Now().Year())
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get cholesterol history record. Error : %v\n", err)
		return nil, err
	}

	activities, err := as.ar.GetAllActivity(c, userID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get all activity. Error : %v\n", err)
		return nil, err
	}

	count := 0
	isHealhtyFoodDone := false
	for _, activity := range activities {
		if activity.IsDone {

			if err := as.ar.DeleteActivity(c, activity.ID); err != nil {
				log.Printf("[cordova-activity-service] failed to delete activity. Error : %v\n", err)
				return nil, err
			}

			count += 1
		}

		if activity.IsDone && activity.Activity == "Healthy Food" {
			isHealhtyFoodDone = true
		}
	}

	if count == 0 {
		return activities, errors.ErrNotYetActivityDone
	}

	reccomendedActivity := activity.GenerateRecommendedActivity(int(userCholesterol.LastCholesterolRecord), count, isHealhtyFoodDone)
	_, err = as.ar.SavedActivity(c, userID, reccomendedActivity)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to save activity. Error : %v\n", err)
		return nil, err
	}

	activities, err = as.ar.GetAllActivity(c, userID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get all activity. Error : %v\n", err)
		return nil, err
	}

	return activities, nil
}

func (as *ActivityService) ChecklistActivity(c context.Context, activityID int) ([]*domain.ActivityDB, error) {
	subActivity, err := as.ar.GetSubActivityByID(c, activityID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get sub activity. Error : %v\n", err)
		return nil, err
	}

	subActivity.IsDone = true

	_, err = as.ar.UpdateSubActivity(c, subActivity)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to update sub activity. Error : %v\n", err)
		return nil, err
	}

	activity, err := as.ar.GetActivityByID(c, subActivity.ActivityID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get activity. Error : %v\n", err)
		return nil, err
	}

	activity.FinishedSubActivity += 1
	if activity.FinishedSubActivity == activity.TotalSubActivity {
		activity.IsDone = true
	}

	_, err = as.ar.UpdateActivity(c, activity)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to update activity. Error : %v\n", err)
		return nil, err
	}

	activities, err := as.ar.GetAllActivity(c, activity.UserID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get all activity. Error : %v\n", err)
		return nil, err
	}

	return activities, nil
}

func (as *ActivityService) UnchecklistActivity(c context.Context, activityID int) ([]*domain.ActivityDB, error) {
	subActivity, err := as.ar.GetSubActivityByID(c, activityID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get sub activity. Error : %v\n", err)
		return nil, err
	}

	subActivity.IsDone = false

	_, err = as.ar.UpdateSubActivity(c, subActivity)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to update sub activity. Error : %v\n", err)
		return nil, err
	}

	activity, err := as.ar.GetActivityByID(c, subActivity.ActivityID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get activity. Error : %v\n", err)
		return nil, err
	}

	activity.FinishedSubActivity -= 1
	if activity.FinishedSubActivity != activity.TotalSubActivity {
		activity.IsDone = false
	}

	_, err = as.ar.UpdateActivity(c, activity)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to update activity. Error : %v\n", err)
		return nil, err
	}

	activities, err := as.ar.GetAllActivity(c, activity.UserID)
	if err != nil {
		log.Printf("[cordova-activity-service] failed to get all activity. Error : %v\n", err)
		return nil, err
	}

	return activities, nil
}
