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
