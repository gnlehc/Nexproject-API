package response

import "github.com/google/uuid"

type SMELoginResponseDTO struct {
	StatusCode int
	Message    string
	Token      string
	Data       SMELoginData
}

type SMELoginData struct {
	SMEID uuid.UUID
	Email string
}
