package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/utils"
)

func Authentication(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "access token not found", errors.New("token are missing"), nil))
	}
	claims, err := utils.ValidateToken(accessToken)

	if err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "invalid token", errors.New(err), nil))
	}
	c.Locals("user", *claims)
	return c.Next()
}
