package repository

import (
	"context"
	"time"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl interface {
	UpdateUser(c context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(c context.Context, val string) (*domain.User, error)
}

type UserRepository struct {
	db *database.ClientDB
}

func NewUserRepository(db *database.ClientDB) UserRepositoryImpl {
	return &UserRepository{db}
}

func (ur *UserRepository) UpdateUser(c context.Context, user *domain.User) (*domain.User, error) {
	argKV := map[string]interface{}{
		"id":                     user.ID,
		"name":                   user.Name,
		"email":                  user.Email,
		"birthday":               user.Birthday,
		"gender":                 user.Gender,
		"weight":                 user.Weight,
		"height":                 user.Height,
		"exercise":               user.Exercise,
		"physical_activity":      user.PhysicalActivity,
		"sleep_hours":            user.SleepHours,
		"smoking":                user.Smoking,
		"alcohol_consumption":    user.AlcoholConsumption,
		"sedentary_hours":        user.SedentaryHours,
		"diabetes":               user.Diabetes,
		"family_history":         user.FamilyHistory,
		"previous_heart_problem": user.PreviousHeartProblem,
		"medication_use":         user.MedicationUse,
		"stress_level":           user.StressLevel,
		"photo_profile":          user.PhotoProfile,
	}

	_, err := ur.db.NamedExecContext(c, UpdateUser, argKV)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(c context.Context, val string) (*domain.User, error) {
	argKV := map[string]interface{}{
		"id": val,
	}

	query, args, err := sqlx.Named(GetAccountByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ur.db.Rebind(query)

	var user UserDB
	if err := ur.db.QueryRowxContext(c, query, args...).StructScan(&user); err != nil {
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
