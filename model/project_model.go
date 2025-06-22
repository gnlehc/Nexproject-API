package model

import (
	"github.com/google/uuid"
)

type Project struct {
	ProjectID          uuid.UUID `json:"ProjectID" gorm:"type:uuid;primaryKey;not null"`
	ProjectName        string    `json:"ProjectName"`
	ProjectDescription string    `json:"ProjectDescription"`
	Jobs               []Job     `json:"Jobs" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SMEID              uuid.UUID `json:"SMEID"`
}
