package user_controller

import (
	"rabi-salon/app_context"
	"rabi-salon/frameworks/database"
	"rabi-salon/usecases/user_case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) Paginate(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("Page", "0"))
	if err != nil {
		return err
	}

	pageSize, err := strconv.Atoi(ctx.Query("PageSize", "10"))
	if err != nil {
		return err
	}

	filter := user_case.PaginateFilter{}
	if err = ctx.QueryParser(&filter); err != nil {
		return err
	}

	paginate := database.PaginateInput{
		Page:     page,
		PageSize: pageSize,
	}

	result, err := c.usecase.Paginate(app_context.New(ctx.Context()), filter, paginate)

	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
