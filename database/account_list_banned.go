package database

import "minetest-skin-server/models"

// AccountListBanned Return users that are banned
func AccountListBanned() ([]models.Account, error) {
	var result []models.Account

	if err := DB.Find(&result).Where("ban = true").Error; err != nil {
		return nil, err
	}

	return result, nil
}
