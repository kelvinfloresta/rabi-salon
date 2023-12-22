package user_gateway

import (
	"rabi-salon/frameworks/database/gorm_adapter/models"
)

func (g *GormUserGatewayAdapter) Delete(id string) (bool, error) {
	result := g.DB.Conn.Where(
		"id = ?", id,
	).Delete(&models.User{})

	return result.RowsAffected > 0, result.Error
}
