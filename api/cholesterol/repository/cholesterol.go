package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CholesterolRepositoryImpl interface {
	SavedRecordCholesterol(c context.Context, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error)
	GetCholesterolHistory(c context.Context, id string) ([]*domain.CholesterolDB, error)
	GetUserByID(c context.Context, id string) (*domain.User, error)
	CountCholesterolRecord(c context.Context, id string, month, year int) (int, error)
	SavedActivity(c context.Context, id string, activity []*domain.Activity) ([]*domain.Activity, error)
	UpdateRecordCholesterol(c context.Context, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error)
	GetAllActivity(c context.Context, id string) ([]*domain.ActivityDB, error)
}

type CholesterolRepository struct {
	db *database.ClientDB
}

func NewCholesterolRepository(db *database.ClientDB) CholesterolRepositoryImpl {
	return &CholesterolRepository{db}
}

func (cr *CholesterolRepository) GetCholesterolHistory(c context.Context, id string) ([]*domain.CholesterolDB, error) {
	var cholesterol []*domain.CholesterolDB

	argKV := map[string]interface{}{
		"user_id": id,
	}

	query, args, err := sqlx.Named(GetCholesterolHistory, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = cr.db.Rebind(query)

	rows, err := cr.db.QueryxContext(c, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cholesterolDB domain.CholesterolDB
		err := rows.StructScan(&cholesterolDB)
		if err != nil {
			return nil, err
		}

		cholesterol = append(cholesterol, &cholesterolDB)
	}

	return cholesterol, nil
}

func (cr *CholesterolRepository) SavedRecordCholesterol(c context.Context, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error) {
	argKV := map[string]interface{}{
		"user_id":                 id,
		"average_cholesterol":     cholesterol.AverageCholesterol,
		"last_cholesterol_record": cholesterol.LastCholesterolRecord,
		"cholesterol_level":       cholesterol.CholesterolLevel,
		"triglycerides":           cholesterol.Triglycerides,
		"heart_rate":              cholesterol.HeartRate,
		"blood_pressure":          cholesterol.BloodPressure,
		"month":                   cholesterol.Month,
		"year":                    cholesterol.Year,
		"heart_risk_percentage":   cholesterol.HeartRiskPercentage,
		"cholesterol_test_date":   cholesterol.CholesterolTestDate,
	}

	_, err := cr.db.NamedExecContext(c, SavedRecordCholesterol, argKV)
	if err != nil {
		return nil, err
	}

	return cholesterol, nil
}

func (cr *CholesterolRepository) UpdateRecordCholesterol(c context.Context, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error) {
	argKV := map[string]interface{}{
		"user_id":                 id,
		"average_cholesterol":     cholesterol.AverageCholesterol,
		"last_cholesterol_record": cholesterol.LastCholesterolRecord,
		"cholesterol_level":       cholesterol.CholesterolLevel,
		"triglycerides":           cholesterol.Triglycerides,
		"heart_rate":              cholesterol.HeartRate,
		"blood_pressure":          cholesterol.BloodPressure,
		"heart_risk_percentage":   cholesterol.HeartRiskPercentage,
		"cholesterol_test_date":   cholesterol.CholesterolTestDate,
		"month":                   cholesterol.Month,
		"year":                    cholesterol.Year,
	}

	_, err := cr.db.NamedExecContext(c, UpdateRecordCholesterol, argKV)
	if err != nil {
		return nil, err
	}

	return cholesterol, nil
}

func (cr *CholesterolRepository) GetUserByID(c context.Context, id string) (*domain.User, error) {
	argKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetUserByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = cr.db.Rebind(query)

	var user UserDB
	if err := cr.db.QueryRowxContext(c, query, args...).StructScan(&user); err != nil {
		return nil, err
	}

	return user.Parse(), nil
}

func (cr *CholesterolRepository) CountCholesterolRecord(c context.Context, id string, month, year int) (int, error) {
	argKV := map[string]interface{}{
		"user_id": id,
		"month":   month,
		"year":    year,
	}

	query, args, err := sqlx.Named(CountCholesterolRecord, argKV)
	if err != nil {
		return -1, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return -1, err
	}
	query = cr.db.Rebind(query)

	var countEmail int
	if err := cr.db.QueryRowxContext(c, query, args...).Scan(&countEmail); err != nil {
		return -1, err
	}

	return countEmail, nil
}

func (cr *CholesterolRepository) SavedActivity(c context.Context, userID string, activity []*domain.Activity) ([]*domain.Activity, error) {
	txClient, err := cr.db.Beginx()
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
			"user_id":               userID,
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
					"description":  value.SubActivities.Description,
					"ingredients":  value.SubActivities.Ingredients,
					"steps":        value.SubActivities.Steps,
					"duration":     value.SubActivities.Duration,
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
					"description":  value.SubActivities.Description,
					"ingredients":  value.SubActivities.Ingredients,
					"steps":        value.SubActivities.Steps,
					"duration":     value.SubActivities.Duration,
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

func (cr *CholesterolRepository) GetAllActivity(c context.Context, id string) ([]*domain.ActivityDB, error) {
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

	query = cr.db.Rebind(query)

	rows, err := cr.db.QueryxContext(c, query, args...)
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

			subActivityID          int
			subActivityActivityID  int
			subActivity            string
			descriptionSubActivity string
			ingredientsSubActivity []string
			stepsSubActivity       []string
			durationSubActivity    int
			isDoneSubActivity      bool
			createdAtSubActivity   time.Time
			updatedAtSubActivity   time.Time
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
			&descriptionSubActivity,
			pq.Array(&ingredientsSubActivity),
			pq.Array(&stepsSubActivity),
			&durationSubActivity,
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
					Ingredients: ingredientsSubActivity,
					Steps:       stepsSubActivity,
					Duration:    durationSubActivity,
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
				Ingredients: ingredientsSubActivity,
				Steps:       stepsSubActivity,
				Duration:    durationSubActivity,
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

type UserDB struct {
	ID                   string     `db:"id"`
	Name                 string     `db:"name"`
	Email                string     `db:"email"`
	Birthday             *time.Time `db:"birthday"`
	Gender               *bool      `db:"gender"`
	Weight               *float64   `db:"weight"`
	Height               *float64   `db:"height"`
	Exercise             *float64   `db:"exercise"`
	PhysicalActivity     *float64   `db:"physical_activity"`
	SleepHours           *float64   `db:"sleep_hours"`
	Smoking              *bool      `db:"smoking"`
	AlcoholConsumption   *bool      `db:"alcohol_consumption"`
	SedentaryHours       *float64   `db:"sedentary_hours"`
	Diabetes             *bool      `db:"diabetes"`
	FamilyHistory        *bool      `db:"family_history"`
	PreviousHeartProblem *bool      `db:"previous_heart_problem"`
	MedicationUse        *bool      `db:"medication_use"`
	StressLevel          *float64   `db:"stress_level"`
	PhotoProfile         *string    `db:"photo_profile"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            time.Time  `db:"updated_at"`
}

func (u *UserDB) Parse() *domain.User {
	user := &domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	if u.Birthday != nil {
		user.Birthday = *u.Birthday
	}

	if u.Gender != nil {
		user.Gender = *u.Gender
	}

	if u.Weight != nil {
		user.Weight = *u.Weight
	}

	if u.Height != nil {
		user.Height = *u.Height
	}

	if u.Exercise != nil {
		user.Exercise = *u.Exercise
	}

	if u.PhysicalActivity != nil {
		user.PhysicalActivity = *u.PhysicalActivity
	}

	if u.SleepHours != nil {
		user.SleepHours = *u.SleepHours
	}

	if u.Smoking != nil {
		user.Smoking = *u.Smoking
	}

	if u.AlcoholConsumption != nil {
		user.AlcoholConsumption = *u.AlcoholConsumption
	}

	if u.SedentaryHours != nil {
		user.SedentaryHours = *u.SedentaryHours
	}

	if u.Diabetes != nil {
		user.Diabetes = *u.Diabetes
	}

	if u.FamilyHistory != nil {
		user.FamilyHistory = *u.FamilyHistory
	}

	if u.PreviousHeartProblem != nil {
		user.PreviousHeartProblem = *u.PreviousHeartProblem
	}

	if u.MedicationUse != nil {
		user.MedicationUse = *u.MedicationUse
	}

	if u.StressLevel != nil {
		user.StressLevel = *u.StressLevel
	}

	if u.PhotoProfile != nil {
		user.PhotoProfile = *u.PhotoProfile
	}

	return user
}
