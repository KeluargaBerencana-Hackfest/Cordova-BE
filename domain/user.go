package domain

type User struct {
	ID                   string  `json:"id"`
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
	Diet                 string  `json:"diet_level"`
	SedentaryHours       float64 `json:"sedentary_hours"`
	HeartRate            float64 `json:"heart_rate"`
	Diabetes             bool    `json:"diabetes"`
	Triglycerides        float64 `json:"triglycerides"`
	FamilyHistory        bool    `json:"family_history"`
	PreviousHeartProblem bool    `json:"previous_heart_problem"`
	MedicationUse        bool    `json:"medication_use"`
}
