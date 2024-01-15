package domain

import "time"

type CholesterolDB struct {
	UserID                string    `json:"-" db:"user_id"`
	AverageCholesterol    float64   `json:"average_cholesterol" db:"average_cholesterol"`
	LastCholesterolRecord float64   `json:"last_cholesterol_record" db:"last_cholesterol_record"`
	CholesterolLevel      string    `json:"cholesterol_level" db:"cholesterol_level"`
	Triglycerides         float64   `json:"triglycerides" db:"triglycerides"`
	HeartRate             float64   `json:"heart_rate" db:"heart_rate"`
	BloodPressure         string    `json:"blood_pressure" db:"blood_pressure"`
	Month                 int       `json:"month" db:"month"`
	Year                  int       `json:"-" db:"year"`
	HeartRiskPercentage   float64   `json:"heart_risk_percentage" db:"heart_risk_percentage"`
	CholesterolTestDate   time.Time `json:"cholesterol_test_date" db:"cholesterol_test_date"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}

type Cholesterol struct {
	UserID       string                   `json:"user_id"`
	Cholesterols map[int][]*CholesterolDB `json:"cholesterols"`
}

type CholesterolRequest struct {
	Triglycerides       float64 `json:"triglycerides"`
	Cholesterol         float64 `json:"cholesterol"`
	BloodPressure       string  `json:"blood_pressure"`
	HeartRate           float64 `json:"heart_rate"`
	CholesterolTestDate string  `json:"cholesterol_test_date"`
}

type CholesterolResponse struct {
	Cholesterol *Cholesterol  `json:"cholesterol"`
	Activity    []*ActivityDB `json:"activity"`
}
