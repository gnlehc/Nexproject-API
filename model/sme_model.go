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
	// Foreignkey
	SMETypeID uuid.UUID `json:"SMETypeID" gorm:"type:uuid"`
	SMEType   SMEType   `json:"SMEType" gorm:"foreignKey:SMETypeID"`
	Projects  []Project `json:"Projects" gorm:"foreignKey:SMEID"`
}
