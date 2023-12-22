package user_gateway

import (
	"rabi-salon/frameworks/database/gorm_adapter/models"
)

func (g *GormUserGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.User{}).Where("id = ?", filter.ID)

	result := query.Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
