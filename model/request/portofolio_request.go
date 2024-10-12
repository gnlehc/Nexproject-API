package request

import "github.com/google/uuid"

type PortofolioRequestDTO struct {
	PortofolioID uuid.UUID
	CoverImage   string
	Title        string
	Description  string
	ProjectLink  string
}
