package domain

import "time"

type User struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Birthday             time.Time `json:"birthday"`
	Gender               bool      `json:"gender"`
	Weight               float64   `json:"weight"`
	Height               float64   `json:"height"`
	Exercise             float64   `json:"exercise"`
	PhysicalActivity     float64   `json:"physical_activity"`
	SleepHours           float64   `json:"sleep_hours"`
	Smoking              bool      `json:"smoking"`
	AlcoholConsumption   bool      `json:"alcohol_consumption"`
	SedentaryHours       float64   `json:"sedentary_hours"`
	Diabetes             bool      `json:"diabetes"`
	FamilyHistory        bool      `json:"family_history"`
	PreviousHeartProblem bool      `json:"previous_heart_problem"`
	MedicationUse        bool      `json:"medication_use"`
	PhotoProfile         string    `json:"photo_profile"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	Name                 string  `json:"name"`
	Email                string  `json:"email"`
	Birthday             string  `json:"birthday"`
	Gender               bool    `json:"gender"`
	Weight               float64 `json:"weight"`
	Height               float64 `json:"height"`
	Exercise             float64 `json:"exercise"`
	PhysicalActivity     float64 `json:"physical_activity"`
	SleepHours           float64 `json:"sleep_hours"`
	Smoking              bool    `json:"smoking"`
	AlcoholConsumption   bool    `json:"alcohol_consumption"`
	SedentaryHours       float64 `json:"sedentary_hours"`
	Diabetes             bool    `json:"diabetes"`
	FamilyHistory        bool    `json:"family_history"`
	PreviousHeartProblem bool    `json:"previous_heart_problem"`
	MedicationUse        bool    `json:"medication_use"`
}
