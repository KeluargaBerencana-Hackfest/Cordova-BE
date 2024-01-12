package repository

import (
	"context"
	"time"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
)

type CholesterolRepositoryImpl interface {
	SavedRecordCholesterol(c context.Context, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error)
	GetCholesterolHistory(c context.Context, id string) ([]*domain.CholesterolDB, error)
	GetUserByID(c context.Context, id string) (*domain.User, error)
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

	rows, err := cr.db.Queryx(query, args...)
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
	if err := cr.db.QueryRowx(query, args...).StructScan(&user); err != nil {
		return nil, err
	}

	return user.Parse(), nil
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
