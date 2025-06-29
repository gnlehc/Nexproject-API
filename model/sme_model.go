package model

import "github.com/google/uuid"

type SME struct {
	SMEID              uuid.UUID `json:"SMEID" gorm:"type:uuid;primaryKey;not null"`
	Email              string    `json:"Email"`
	Password           string    `json:"Password"`
	CompanyName        string    `json:"CompanyName"`
	CompanyDescription string    `json:"CompanyDescription"`
	CEO                string    `json:"CEO"`
	Social             string    `json:"Social"`
	PhoneNumber        string    `json:"PhoneNumber"`
	ActiveStatus       bool      `json:"ActiveStatus"`
	SMEType            string    `json:"SMEType"`
	// Foreignkey
	Projects []Project `json:"Projects" gorm:"foreignKey:SMEID"`
	Location string    `json:"Location"`
}
