package domain

import "time"

type SubActivityDB struct {
	ID          int       `json:"id"`
	ActivityID  int       `json:"activity_id"`
	SubActivity string    `json:"sub_activity"`
	Description string    `json:"description"`
	Ingredients []string  `json:"ingredients"`
	Steps       []string  `json:"steps"`
	Duration    int       `json:"duration"`
	IsDone      bool      `json:"is_done"`
	Image       string    `json:"image" db:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActivityDB struct {
	ID                  int                        `json:"id" db:"id"`
	UserID              string                     `json:"user_id" db:"user_id"`
	Activity            string                     `json:"activity" db:"activity"`
	Description         string                     `json:"description" db:"description"`
	TotalSubActivity    int                        `json:"total_sub_activity" db:"total_sub_activity"`
	FinishedSubActivity int                        `json:"finished_sub_activity" db:"finished_sub_activity"`
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
	Duration        int //minutes
	IsSequential    bool
	Count           int
	Image           string
}

type Activity struct {
	ID            int
	NameActivity  string
	Description   string
	Duration      int
	SubActivities SubActivity
}
