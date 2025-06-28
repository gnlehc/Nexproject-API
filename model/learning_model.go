package model

import "github.com/google/uuid"

type Learning struct {
	LearningID    uuid.UUID `json:"LearningID" gorm:"type:uuid;primaryKey;not null"`
	Title         string    `json:"Title"`
	Content       string    `json:"Content"`
	SkillID       uuid.UUID `json:"SkillID"`
	Skill         Skill
	ImageCoverURL string `json:"ImageCoverURL"`
}
