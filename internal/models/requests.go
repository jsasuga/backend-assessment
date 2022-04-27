package models

import "time"

type RequestNewWorkout struct {
	Athlete   string    `json:"athlete"`
	Coach     string    `json:"coach"`
	Scheduled time.Time `json:"scheduled"`
}

type UpdateWorkout struct {
	Scheduled   time.Time `json:"scheduled"`
	Description string    `json:"description"`
}
