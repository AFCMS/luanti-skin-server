package database

import "luanti-skin-server/backend/models"

// SkinRecent Return the `count` recently published skins
func SkinRecent(count uint8) ([]models.Skin, error) {
	var results []models.Skin

	if err := DB.Find(&results).Where("public = true").Order("created_at DESC").Limit(int(count)).Error; err != nil {
		return results, err
	}
	return results, nil
}
