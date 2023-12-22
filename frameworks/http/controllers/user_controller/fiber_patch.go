package user_controller

import (
	"rabi-salon/app_context"
	"rabi-salon/frameworks/http/fiber_adapter/parser"
	"rabi-salon/usecases/user_case"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) Patch(ctx *fiber.Ctx) error {
	filter := &user_case.PatchFilter{
		ID: ctx.Params("id"),
	}

	data := user_case.PatchValues{}
	if err := parser.ParseBody(ctx, &data); err != nil {
		return ctx.JSON(err)
	}

	updated, err := c.usecase.Patch(app_context.New(ctx.Context()), *filter, data)

	if err != nil {
		return err
	}

	if updated {
		return ctx.SendStatus(200)
	}

	return ctx.SendStatus(404)
}
