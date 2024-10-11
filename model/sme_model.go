package model

import "github.com/google/uuid"

type SME struct {
	SMEID              uuid.UUID `json:"SMEID" gorm:"primaryKey;autoIncrement:true;not null"`
	Email              string    `json:"Email"`
	Password           string    `json:"Password"`
	CompanyName        string    `json:"CompanyName"`
	CompanyDescription string    `json:"CompanyDescription"`
	CEO                string    `json:"CEO"`
	Social             string    `json:"Social"`
	PhoneNumber        string    `json:"PhoneNumber"`
	SMEType            string    `json:"SMEType"`
	ActiveStatus       bool      `json:"ActiveStatus"`
	JobsOpening        []Job     `json:"JobsOpening" gorm:"type:json"`
}
