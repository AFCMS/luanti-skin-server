package database

import "luanti-skin-server/backend/models"

// AccountList Return users that are not banned
func AccountList() ([]models.Account, error) {
	var result []models.Account

	if err := DB.Find(&result).Where("ban = false").Error; err != nil {
		return nil, err
	}

	return result, nil
}
