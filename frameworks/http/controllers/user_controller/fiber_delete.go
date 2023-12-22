package user_controller

import (
	"rabi-salon/app_context"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	deleted, err := c.usecase.Delete(app_context.New(ctx.Context()), id)

	if err != nil {
		return err
	}

	if deleted {
		return ctx.SendStatus(204)
	}

	return ctx.SendStatus(404)
}
