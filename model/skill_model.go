package model

import "github.com/google/uuid"

type Skill struct {
	SkillID   uuid.UUID  `json:"SkillID" gorm:"type:uuid;primaryKey;not null"`
	SkillName string     `json:"SkillName"`
	Talents   []Talent   `json:"Talents" gorm:"many2many:talent_skills"`
	Jobs      []Job      `json:"Jobs" gorm:"many2many:job_skills"`
	Learnings []Learning `json:"Learnings" gorm:"many2many:learning_skills"`
}
