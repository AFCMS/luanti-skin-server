package database

import (
	"luanti-skin-server/backend/models"

	"github.com/google/uuid"
)

// SkinFromUUID Get skin from UUID
func SkinFromUUID(id uuid.UUID) (models.Skin, error) {
	var s = models.Skin{}

	if err := DB.Where("uuid = ?", id).First(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}
