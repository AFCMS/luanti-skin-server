package database

import "luanti-skin-server/backend/models"

// AccountFromUsername Get account from name
func AccountFromUsername(username string) (models.Account, error) {
	var a = models.Account{}

	if err := DB.Where("username = ?", username).First(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
