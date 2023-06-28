package middleware

import (
	"github.com/RianNegreiros/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if !user.IsAdmin {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
