package model

import "github.com/google/uuid"

type Skill struct {
	SkillID   uuid.UUID `json:"SkillID" gorm:"primaryKey;autoIncrement:true;not null"`
	SkillName string    `json:"SkillName"`
}
