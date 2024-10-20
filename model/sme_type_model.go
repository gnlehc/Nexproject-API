package model

import "github.com/google/uuid"

type SMEType struct {
	SMETypeID          uuid.UUID `json:"SMETypeID" gorm:"type:uuid;primaryKey;not null"`
	SMETypeName        string    `json:"SMETypeName"`
	SMETypeDescription string    `json:"SMETypeDescription"`
}
