package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skin struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid"`
	Description string    `json:"description" gorm:"not null"`
	OwnerID     uint      `json:"owner_id"`
	Owner       Account   `json:"-" gorm:"not null"`
	Data        []byte    `json:"-" gorm:"not null"`
	DataHead    []byte    `json:"-" gorm:"not null"`
	CreatedAt   time.Time `json:"creation_date" gorm:"not null"`
}

func (base *Skin) BeforeCreate(tx *gorm.DB) error {
	base.UUID = uuid.New()

	return nil
}
