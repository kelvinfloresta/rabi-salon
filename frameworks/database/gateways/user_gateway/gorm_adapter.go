package user_gateway

import "rabi-salon/frameworks/database/gorm_adapter"

type GormUserGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
