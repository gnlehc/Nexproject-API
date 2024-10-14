package model

import "github.com/google/uuid"

type TrApplication struct {
	AppID uuid.UUID `json:"AppID" gorm:"primaryKey;autoIncrement:true;not null"`
	// Foreignkey
	TalentID uuid.UUID `json:"TalentID" gorm:"type:uuid"`
	Talent   Talent    `json:"Talent" gorm:"foreignKey:TalentID"`
	JobID    uuid.UUID `json:"JobID" gorm:"type:uuid"`
	Job      Job       `json:"Job" gorm:"foreignKey:JobID"`
	Status   string    `json:"Status"`
	Rating   float32   `json:"Rating"`
	Comment  string    `json:"Comment" gorm:"type:text"`
}
