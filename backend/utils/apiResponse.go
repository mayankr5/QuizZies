package utils

import "github.com/gofiber/fiber"

func APIResponse(status string, msg string, err error, data any) fiber.Map {
	var errRes any
	if err != nil {
		errRes = err.Error()
	} else {
		errRes = nil
	}

	return fiber.Map{
		"status":  status,
		"message": msg,
		"data":    data,
		"error":   errRes,
	}
}
