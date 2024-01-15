package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
)

type ActivityRepositoryImpl interface {
	UpdateSubActivity(c context.Context, subActivity *domain.SubActivityDB) (*domain.SubActivityDB, error)
	UpdateActivity(c context.Context, activity *domain.ActivityDB) (*domain.ActivityDB, error)
	GetAllActivity(c context.Context, id string) ([]*domain.ActivityDB, error)
	GetSubActivityByID(c context.Context, subActivityID int) (*domain.SubActivityDB, error)
	GetActivityByID(c context.Context, activityID int) (*domain.ActivityDB, error)
	GetUserCholesterolByID(c context.Context, userID string, month, year int) (*domain.CholesterolDB, error)
	SavedActivity(c context.Context, id string, activity []*domain.Activity) ([]*domain.Activity, error)
	DeleteActivity(c context.Context, activityID int) error
}

type ActivityRepository struct {
	db *database.ClientDB
}

func NewActivitylRepository(db *database.ClientDB) ActivityRepositoryImpl {
	return &ActivityRepository{db}
}

func (ar *ActivityRepository) GetAllActivity(c context.Context, id string) ([]*domain.ActivityDB, error) {
	argKV := map[string]interface{}{
		"user_id": id,
	}

	query, args, err := sqlx.Named(GetAllActivity, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ar.db.Rebind(query)
	rows, err := ar.db.QueryxContext(c, query, args...)
	if err != nil {
		return nil, err
	}

	var activities []*domain.ActivityDB

	for rows.Next() {
		var (
			activityID          int
			userID              string
			activity            string
			description         string
			totalSubActivity    int
			finishedSubActivity int
			image               string
			isDoneActivity      bool
			createdAtActivity   time.Time
			updatedAtActivity   time.Time

			subActivityID         int
			subActivityActivityID int
			subActivity           string
			isDoneSubActivity     bool
			createdAtSubActivity  time.Time
			updatedAtSubActivity  time.Time
		)

		if err := rows.Scan(
			&activityID,
			&userID,
			&activity,
			&description,
			&totalSubActivity,
			&finishedSubActivity,
			&image,
			&isDoneActivity,
			&createdAtActivity,
			&updatedAtActivity,
			&subActivityID,
			&subActivityActivityID,
			&subActivity,
			&isDoneSubActivity,
			&createdAtSubActivity,
			&updatedAtSubActivity,
		); err != nil {
			return nil, err
		}

		var found bool
		for i := range activities {
			if activities[i].ID == activityID {
				subActivityDB := domain.SubActivityDB{
					ID:          subActivityID,
					ActivityID:  subActivityActivityID,
					SubActivity: subActivity,
					IsDone:      isDoneSubActivity,
					CreatedAt:   createdAtSubActivity,
					UpdatedAt:   updatedAtSubActivity,
				}
				activities[i].SubActivities[activity] = append(activities[i].SubActivities[activity], subActivityDB)
				found = true
				break
			}
		}

		if !found {
			activityDB := domain.ActivityDB{
				ID:                  activityID,
				UserID:              userID,
				Activity:            activity,
				Description:         description,
				TotalSubActivity:    totalSubActivity,
				FinishedSubActivity: finishedSubActivity,
				Image:               image,
				IsDone:              isDoneActivity,
				CreatedAt:           createdAtActivity,
				UpdatedAt:           updatedAtActivity,
				SubActivities:       make(map[string][]domain.SubActivityDB),
			}

			subActivityDB := domain.SubActivityDB{
				ID:          subActivityID,
				ActivityID:  subActivityActivityID,
				SubActivity: subActivity,
				IsDone:      isDoneSubActivity,
				CreatedAt:   createdAtSubActivity,
				UpdatedAt:   updatedAtSubActivity,
			}

			activityDB.SubActivities[activity] = append(activityDB.SubActivities[activity], subActivityDB)
			activities = append(activities, &activityDB)
		}
	}

	return activities, nil
}

func (ar *ActivityRepository) UpdateSubActivity(c context.Context, subActivity *domain.SubActivityDB) (*domain.SubActivityDB, error) {
	argKV := map[string]interface{}{
		"id":           subActivity.ID,
		"sub_activity": subActivity.SubActivity,
		"ingredients":  subActivity.Ingredients,
		"steps":        subActivity.Steps,
		"is_done":      subActivity.IsDone,
	}

	_, err := ar.db.NamedExecContext(c, UpdateActivity, argKV)
	if err != nil {
		return nil, err
	}

	return subActivity, nil
}

func (ar *ActivityRepository) GetSubActivityByID(c context.Context, subActivityID int) (*domain.SubActivityDB, error) {
	argKV := map[string]interface{}{
		"id": subActivityID,
	}

	query, args, err := sqlx.Named(GetSubActivityByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ar.db.Rebind(query)

	var subActivity *domain.SubActivityDB
	if err := ar.db.QueryRowxContext(c, query, args...).StructScan(&subActivity); err != nil {
		return nil, err
	}

	return subActivity, nil
}

func (ar *ActivityRepository) UpdateActivity(c context.Context, activity *domain.ActivityDB) (*domain.ActivityDB, error) {
	argKV := map[string]interface{}{
		"id":                    activity.ID,
		"activity":              activity.Activity,
		"description":           activity.Description,
		"image":                 activity.Image,
		"is_done":               activity.IsDone,
		"total_sub_activity":    activity.TotalSubActivity,
		"finished_sub_activity": activity.FinishedSubActivity,
	}

	_, err := ar.db.NamedExecContext(c, UpdateActivity, argKV)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (ar *ActivityRepository) GetActivityByID(c context.Context, activityID int) (*domain.ActivityDB, error) {
	argKV := map[string]interface{}{
		"id": activityID,
	}

	query, args, err := sqlx.Named(GetActivityByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ar.db.Rebind(query)

	var subActivity *domain.ActivityDB
	if err := ar.db.QueryRowxContext(c, query, args...).StructScan(&subActivity); err != nil {
		return nil, err
	}

	return subActivity, nil
}

func (ar *ActivityRepository) GetUserCholesterolByID(c context.Context, userID string, month, year int) (*domain.CholesterolDB, error) {
	argKV := map[string]interface{}{
		"id":    userID,
		"month": month,
		"year":  year,
	}

	query, args, err := sqlx.Named(GetUserByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ar.db.Rebind(query)

	var cholesterol *domain.CholesterolDB
	if err := ar.db.QueryRowx(query, args...).StructScan(&cholesterol); err != nil {
		return nil, err
	}

	return cholesterol, nil
}

func (ar *ActivityRepository) SavedActivity(c context.Context, id string, activity []*domain.Activity) ([]*domain.Activity, error) {
	txClient, err := ar.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() error {
		if err != nil {
			txClient.Rollback()
			log.Printf("[cordova-cholesterol-repository] action got rollback, error occured: %s\n", err.Error())
			return err
		}
		return nil
	}()

	for _, value := range activity {
		argKV := map[string]interface{}{
			"user_id":               id,
			"activity":              value.NameActivity,
			"description":           value.Description,
			"total_sub_activity":    value.SubActivities.Count,
			"finished_sub_activity": 0,
			"image":                 value.Image,
			"is_done":               false,
		}

		var activityID int

		query, args, err := sqlx.Named(SavedActivity, argKV)
		if err != nil {
			return nil, err
		}

		query, args, err = sqlx.In(query, args...)
		if err != nil {
			return nil, err
		}

		query = txClient.Rebind(query)

		err = txClient.QueryRowxContext(c, query, args...).Scan(&activityID)
		if err != nil {
			return nil, err
		}

		value.ID = activityID

		if value.SubActivities.IsSequential {
			for i := 0; i < value.SubActivities.Count; i++ {
				argKV := map[string]interface{}{
					"activity_id":  activityID,
					"sub_activity": fmt.Sprintf("%s %d", value.SubActivities.NameSubActivity, i+1),
					"ingredients":  value.SubActivities.Ingredients,
					"steps":        value.SubActivities.Steps,
					"is_done":      false,
				}

				var subActivityID int
				query, args, err := sqlx.Named(SavedSubActivity, argKV)
				if err != nil {
					return nil, err
				}

				query, args, err = sqlx.In(query, args...)
				if err != nil {
					return nil, err
				}

				query = txClient.Rebind(query)
				err = txClient.QueryRowxContext(c, query, args...).Scan(&subActivityID)
				if err != nil {
					return nil, err
				}
				value.SubActivities.ID = subActivityID
			}
		} else {
			for i := 0; i < value.SubActivities.Count; i++ {
				argKV := map[string]interface{}{
					"activity_id":  activityID,
					"sub_activity": value.SubActivities.NameSubActivity,
					"is_done":      false,
				}

				var subActivityID int
				query, args, err := sqlx.Named(SavedSubActivity, argKV)
				if err != nil {
					return nil, err
				}

				query, args, err = sqlx.In(query, args...)
				if err != nil {
					return nil, err
				}

				query = txClient.Rebind(query)
				err = txClient.QueryRowxContext(c, query, args...).Scan(&subActivityID)
				if err != nil {
					return nil, err
				}
				value.SubActivities.ID = subActivityID
			}
		}
	}

	if err = txClient.Commit(); err != nil {
		log.Printf("[cordova-cholesterol-repository] failed to commit the transaction.: %s\n", err.Error())
		return nil, err
	}

	return activity, nil
}

func (ar *ActivityRepository) DeleteActivity(c context.Context, activityID int) error {
	argKV := map[string]interface{}{
		"id": activityID,
	}

	_, err := ar.db.NamedExecContext(c, DeleteActivity, argKV)
	if err != nil {
		return err
	}

	return nil
}
