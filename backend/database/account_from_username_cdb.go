package database

import "luanti-skin-server/backend/models"

// AccountFromCDBUsername Get account from CDB username
func AccountFromCDBUsername(cdbUsername string) (models.Account, error) {
	var a = models.Account{}

	if err := DB.Where("cdb_username = ?", cdbUsername).First(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
