package model

import "github.com/google/uuid"

type Learning struct {
	LearningID    uuid.UUID `json:"LearningID" gorm:"type:uuid;primaryKey;not null"`
	Title         string    `json:"Title"`
	Content       string    `json:"Content"`
	Skills        []Skill   `json:"Skills" gorm:"many2many:learning_skills"` 
	ImageCoverURL string    `json:"ImageCoverURL"`
}
