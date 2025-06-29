package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	JobID          uuid.UUID `json:"JobID" gorm:"type:uuid;primaryKey;not null"`
	JobTitle       string    `json:"JobTitle"`
	JobDescription string    `json:"JobDescription"`
	JobType        string    `json:"JobType"`
	Qualification  string    `json:"Qualification"`
	JobArrangement string    `json:"JobArrangement"`
	Wage           string    `json:"Wage"`
	Active         bool      `json:"Active"`
	CreatedAt      time.Time `json:"CreatedAt"`
	Location       string    `json:"Location"`
	// ForeignKey
	Skills    []Skill   `gorm:"many2many:job_skills"`
	ProjectID uuid.UUID `json:"ProjectID" gorm:"type:uuid;not null"`
}
