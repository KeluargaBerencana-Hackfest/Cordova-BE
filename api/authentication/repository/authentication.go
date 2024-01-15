package repository

import (
	"time"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type AuthRepositoryImpl interface {
	SaveUser(c context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(c context.Context, vas string) (*domain.User, error)
	CountEmailAccount(c context.Context, email string) (int, error)
}

type AuthRepository struct {
	db *database.ClientDB
}

func NewAuthRepository(db *database.ClientDB) AuthRepositoryImpl {
	return &AuthRepository{db}
}

func (ar *AuthRepository) SaveUser(c context.Context, user *domain.User) (*domain.User, error) {
	argKV := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	_, err := ar.db.NamedExecContext(c, SavedAccount, argKV)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ar *AuthRepository) GetUserByID(c context.Context, userID string) (*domain.User, error) {
	argKV := map[string]interface{}{
		"id": userID,
	}

	query, args, err := sqlx.Named(GetAccountByID, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = ar.db.Rebind(query)

	var user UserDB
	if err := ar.db.QueryRowxContext(c, query, args...).StructScan(&user); err != nil {
		return nil, err
	}

	return user.Parse(), nil
}

func (ar *AuthRepository) CountEmailAccount(c context.Context, email string) (int, error) {
	argKV := map[string]interface{}{
		"email": email,
	}

	query, args, err := sqlx.Named(CountEmail, argKV)
	if err != nil {
		return -1, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return -1, err
	}
	query = ar.db.Rebind(query)

	var countEmail int
	if err := ar.db.QueryRowxContext(c, query, args...).Scan(&countEmail); err != nil {
		return -1, err
	}

	return countEmail, nil
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
