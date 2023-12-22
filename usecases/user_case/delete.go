package user_case

import "rabi-salon/app_context"

func (c UserCase) Delete(ctx *app_context.AppContext, id string) (bool, error) {
	return c.gateway.Delete(id)
}
