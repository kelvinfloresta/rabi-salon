package auth_controller

import (
	"errors"
	"fmt"
	"rabi-salon/app_context"
	"rabi-salon/usecases/auth_case/role"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Session(c *fiber.Ctx) error {
	token, ok := c.Context().UserValue("user").(*jwt.Token)
	if !ok || !token.Valid {
		return errors.New("INVALID_TOKEN")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("INVALID_CLAIMS")
	}

	session := &app_context.UserSession{
		UserID:         fmt.Sprint(claims["user_id"]),
		Name:           fmt.Sprint(claims["name"]),
		Login:          fmt.Sprint(claims["login"]),
		OriginalUserID: fmt.Sprint(claims["original_user_id"]),
		Role:           role.Role(fmt.Sprint(claims["role"])),
	}

	c.Context().SetUserValue(app_context.SessionKey, session)

	return c.Next()
}
