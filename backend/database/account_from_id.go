package database

import "luanti-skin-server/backend/models"

// AccountFromID Get account from ID
func AccountFromID(id uint) (models.Account, error) {
	var a = models.Account{}

	if err := DB.Where("id = ?", id).First(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
