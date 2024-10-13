package model

import "github.com/google/uuid"

type Talent struct {
	TalentID       uuid.UUID `json:"TalentID" gorm:"primaryKey;autoIncrement:true;not null"`
	Email          string    `json:"Email"`
	Password       string    `json:"Password"`
	FullName       string    `json:"FullName"`
	Bio            string    `json:"Bio"`
	PhoneNumber    string    `json:"PhoneNumber"`
	ProfilePicture string    `json:"ProfilePicture"`
	ActiveStatus   bool      `json:"ActiveStatus"`
	// Foreignkey
	Skills    []Skill `json:"Skills" gorm:"many2many:talent_skills"`
	AvgRating float32 `json:"AvgRating"`
	HireCount int     `json:"HireCount"`
	// Foreignkey
	Portofolio []Portofolio `json:"Portofolio" gorm:"foreignKey:TalentID"`
	CV         string       `json:"CV"`
	Location   string       `json:"Location"`
}
