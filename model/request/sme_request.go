package request

import "github.com/google/uuid"

type SMELoginRequestDTO struct {
	Email    string
	Password string
}

type SMERegisterRequestDTO struct {
	CompanyName string
	Email       string
	Password    string
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
	SMEType            string    `json:"sme_type"`
	Location           string    `json:"location"`
}
