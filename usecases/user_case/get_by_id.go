package user_case

import (
	"rabi-salon/app_context"
	g "rabi-salon/frameworks/database/gateways/user_gateway"
)

func (c UserCase) GetByID(ctx *app_context.AppContext, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
