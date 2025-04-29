package database

import (
	"luanti-skin-server/backend/models"
)

func AccountSetPermission(a *models.Account, level int8) error {
	if err := DB.Model(&a).Update("permission_level", level).Error; err != nil {
		return err
	}
	return nil
}
