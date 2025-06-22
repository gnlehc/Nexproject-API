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
	University     string    `json:"University"`
	ActiveStatus   bool      `json:"ActiveStatus"`
	AvgRating      float32   `json:"AvgRating"`
	HireCount      int       `json:"HireCount"`
	CV             string    `json:"CV"`
	Location       string    `json:"Location"`
	// Foreignkey
	Skills     []Skill      `json:"Skills" gorm:"many2many:talent_skills"`
	Portofolio []Portofolio `json:"Portofolio" gorm:"foreignKey:TalentID"`
}
