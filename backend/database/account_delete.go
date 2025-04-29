package database

import "luanti-skin-server/backend/models"

// AccountDelete Delete account
func AccountDelete(name string) error {
	return DB.Where("name = ?", name).Delete(&models.Account{}).Error
}
