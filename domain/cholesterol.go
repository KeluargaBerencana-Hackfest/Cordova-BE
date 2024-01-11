package domain

import "time"

type Cholesterol struct {
	UserID           string  `json:"user_id"`
	Cholesterol      float64 `json:"cholesterol"`
	Triglycerides    float64 `json:"triglycerides"`
	CholesterolLevel string  `json:"cholesterol_level"`
	Year             uint64  `json:"year"`
	Month            string  `json:"month"`
}

type CholesterolRequest struct {
	Triglycerides       float64   `json:"triglycerides"`
	Cholesterol         float64   `json:"cholesterol"`
	BloodPressure       string    `json:"blood_pressure"`
	HeartRate           float64   `json:"heart_rate"`
	CholesterolTestDate time.Time `json:"cholesterol_test_date"`
}
