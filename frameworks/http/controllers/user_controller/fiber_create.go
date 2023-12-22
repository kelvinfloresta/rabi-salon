package user_controller

import (
	"rabi-salon/app_context"
	"rabi-salon/usecases/user_case"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) Create(ctx *fiber.Ctx) error {
	data := &user_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(app_context.New(ctx.Context()), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
