package user_controller

import (
	"rabi-salon/app_context"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.usecase.GetByID(app_context.New(ctx.Context()), id)

	if err != nil {
		return err
	}

	if data == nil {
		return ctx.SendStatus(404)
	}

	return ctx.JSON(data)
}
