package models

import "time"

type Workout struct {
	ID          int       `json:"id"`
	Athlete     string    `json:"athlete"`     // email
	Coach       string    `json:"coach"`       // email
	Submitted   time.Time `json:"submitted"`   // time created
	Scheduled   time.Time `json:"scheduled"`   // workout schedule
	Approved    bool      `json:"approved"`    // approved by both
	Completed   bool      `json:"completed"`   // workout completed?
	Description string    `json:"description"` // small description
}
