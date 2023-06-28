package api

import (
	"github.com/RianNegreiros/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}
	return user, nil
}
