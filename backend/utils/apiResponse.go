package utils

import "github.com/gofiber/fiber"

func APIResponse(status string, msg string, err error, data any) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": msg,
		"data":    data,
		"error":   err.Error(),
	}
}
