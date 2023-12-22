package user_case

import (
	"rabi-salon/app_context"
	"rabi-salon/frameworks/database"
	g "rabi-salon/frameworks/database/gateways/user_gateway"
)

type PaginateFilter struct {
	State *string
	City  *string
	Name  *string
}

func (c UserCase) Paginate(ctx *app_context.AppContext, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		City:  input.City,
		State: input.State,
		Name:  input.Name,
	}, paginate)
}
