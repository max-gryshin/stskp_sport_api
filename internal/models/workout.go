package models

import "time"

type Workouts []Workout

type Workout struct {
	ID          int       `json:"id"`
	User        User      `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}
