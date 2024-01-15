package domain

import "time"

type SubActivityDB struct {
	ID          int       `json:"id" db:"id"`
	ActivityID  int       `json:"activity_id" db:"activity_id"`
	SubActivity string    `json:"sub_activity" db:"sub_activity"`
	Ingredients []string  `json:"ingredients" db:"ingredients"`
	Steps       []string  `json:"steps" db:"steps"`
	IsDone      bool      `json:"is_done" db:"is_done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ActivityDB struct {
	ID                  int                        `json:"id" db:"id"`
	UserID              string                     `json:"user_id" db:"user_id"`
	Activity            string                     `json:"activity" db:"activity"`
	Description         string                     `json:"description" db:"description"`
	TotalSubActivity    int                        `json:"total_sub_activity" db:"total_sub_activity"`
	FinishedSubActivity int                        `json:"finished_sub_activity" db:"finished_sub_activity"`
	Image               string                     `json:"image" db:"image"`
	IsDone              bool                       `json:"is_done" db:"is_done"`
	SubActivities       map[string][]SubActivityDB `json:"sub_activities"`
	CreatedAt           time.Time                  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time                  `json:"updated_at" db:"updated_at"`
}

type SubActivity struct {
	ID              int
	NameSubActivity string
	Description     string
	Ingredients     []string
	Steps           []string
	IsSequential    bool
	Count           int
}

type Activity struct {
	ID            int
	NameActivity  string
	Description   string
	Image         string
	Duration      int
	SubActivities SubActivity
}
