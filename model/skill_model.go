package model

import "github.com/google/uuid"

type Skill struct {
	SkillID   uuid.UUID `json:"SkillID" gorm:"primaryKey;not null"`
	Talents   []Talent  `json:"Talents" gorm:"many2many:talent_skills"`
	Jobs      []Job     `json:"Jobs" gorm:"many2many:job_skills"`
	SkillName string    `json:"SkillName"`
}
