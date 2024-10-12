package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	JobID          uuid.UUID `json:"JobID" gorm:"primaryKey;autoIncrement:true;not null"`
	SMEID          uuid.UUID `json:"SMEID" gorm:"type:uuid"`
	SME            SME       `json:"SME" gorm:"foreignKey:SMEID"`
	JobTitle       string    `json:"JobTitle"`
	JobDescription string    `json:"JobDescription"`
	JobType        string    `json:"JobType"`
	Qualification  string    `json:"Qualification"`
	JobArrangement string    `json:"JobArrangement"`
	Wage           string    `json:"Wage"`
	Active         bool      `json:"Active"`
	CreatedAt      time.Time `json:"CreatedAt"`
	Location       string    `json:"Location"`
	Skills         []Skill   `json:"Skills" gorm:"type:json"`
}
