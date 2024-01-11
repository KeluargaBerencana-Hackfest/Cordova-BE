package domain

import "time"

type CholesterolDB struct {
	UserID           string    `json:"-" db:"user_id"`
	Cholesterol      float64   `json:"cholesterol" db:"cholesterol"`
	CholesterolLevel string    `json:"cholesterol_level" db:"cholesterol_level"`
	Triglycerides    float64   `json:"triglycerides" db:"triglycerides"`
	HeartRate        float64   `json:"heart_rate" db:"heart_rate"`
	BloodPressure    string    `json:"blood_pressure" db:"blood_pressure"`
	Month            uint64    `json:"month" db:"month"`
	Year             uint64    `json:"-" db:"year"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type Cholesterol struct {
	UserID       string                      `json:"user_id"`
	Cholesterols map[uint64][]*CholesterolDB `json:"cholesterols"`
}

type CholesterolRequest struct {
	Triglycerides       float64   `json:"triglycerides"`
	Cholesterol         float64   `json:"cholesterol"`
	BloodPressure       string    `json:"blood_pressure"`
	HeartRate           float64   `json:"heart_rate"`
	CholesterolTestDate time.Time `json:"cholesterol_test_date"`
}
