package request

import "github.com/google/uuid"

type SMELoginRequestDTO struct {
	Email    string
	Password string
}

type SMERegisterRequestDTO struct {
	Email       string
	Password    string
	CompanyName string
	CEO         string
	PhoneNumber string
	SMETypeID   uuid.UUID
}

type EditSMERequestDTO struct {
	SMEID              uuid.UUID `json:"sme_id" binding:"required"`
	Email              string    `json:"email"`
	Password           string    `json:"password,omitempty"`
	CompanyName        string    `json:"company_name"`
	CompanyDescription string    `json:"company_description"`
	CEO                string    `json:"ceo"`
	Social             string    `json:"social"`
	PhoneNumber        string    `json:"phone_number"`
	ActiveStatus       bool      `json:"active_status"`
	SMETypeID          uuid.UUID `json:"sme_type_id" gorm:"type:uuid"`
}
