package model

import "github.com/google/uuid"

type SMEType struct {
	SMETypeID          uuid.UUID `json:"SMETypeID" gorm:"primaryKey;autoIncrement:true;not null"`
	SMETypeName        string    `json:"SMETypeName"`
	SMETypeDescription string    `json:"SMETypeDescription"`
}
