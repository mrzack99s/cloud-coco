package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UUID      string         `json:"uuid" gorm:"size:150;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	o.UUID = uuid.New().String()
	return
}
