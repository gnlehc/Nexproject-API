package model

import "github.com/google/uuid"

type SavedJobs struct {
	TalentID uuid.UUID `json:"TalentID" gorm:"type:uuid;not null"` // Make sure this is not nullable if required
	Talent   Talent    `json:"Talent" gorm:"foreignKey:TalentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	JobID uuid.UUID `json:"JobID" gorm:"type:uuid;not null"` // Same here for JobID
	Job   Job       `json:"Job" gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
