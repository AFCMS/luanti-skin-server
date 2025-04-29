package database

import "luanti-skin-server/backend/models"

// SkinCount Count skins in database
func SkinCount() (int64, error) {
	var count int64 = 0
	if err := DB.Model(&models.Skin{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
