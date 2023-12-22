package factories

import (
	"rabi-salon/frameworks/database"
	"rabi-salon/frameworks/database/gateways/user_gateway"
	"rabi-salon/frameworks/database/gorm_adapter"
	"rabi-salon/frameworks/http/controllers/user_controller"

	"rabi-salon/usecases/user_case"
)

func NewUser(d database.Database) *user_controller.UserController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &user_gateway.GormUserGatewayAdapter{DB: DB}
	usecase := user_case.New(gateway)
	return user_controller.New(usecase)
}
