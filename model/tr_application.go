package model

import "github.com/google/uuid"

type TrApplication struct {
	AppID uuid.UUID `json:"AppID" gorm:"primaryKey;autoIncrement:true;not null"`
	// Foreignkey
	TalentID  uuid.UUID         `json:"TalentID" gorm:"type:uuid"`
	Talent    Talent            `json:"Talent" gorm:"foreignKey:TalentID"`
	ProjectID uuid.UUID         `json:"ProjectID" gorm:"type:uuid"`
	Project   Project           `json:"Project" gorm:"foreignKey:JobID"`
	JobID     uuid.UUID         `json:"JobID" gorm:"type:uuid"`
	Job       Job               `json:"Job" gorm:"foreignKey:JobID"`
	StatusID  int               `json:"StatusID"`
	Status    ApplicationStatus `json:"Status" gorm:"foreignKey:StatusID"`
	Rating    float32           `json:"Rating"`
	Comment   string            `json:"Comment" gorm:"type:text"`
}

type ApplicationStatus struct {
	StatusID int    `json:"StatusID" gorm:"primaryKey;autoIncrement:true;not null"`
	Status   string `json:"Status"`
}
