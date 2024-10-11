package model

import "github.com/google/uuid"

type Portofolio struct {
	PortofolioID uuid.UUID `json:"PortofolioID" gorm:"primaryKey;autoIncrement:true;not null"`
	CoverImage   string    `json:"CoverImage"`
	Title        string    `json:"Title"`
	Description  string    `json:"Description"`
	ProjectLink  string    `json:"ProjectLink"`
}
