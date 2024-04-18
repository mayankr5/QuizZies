package handlers

import "github.com/gofiber/fiber/v2"

func GetAllQuestion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "questions fetched",
		"data":    nil,
	})
}

func GetQuestion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "question fetched",
		"data":    nil,
	})
}

func AddQuestion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "question added",
		"data":    nil,
	})
}

func UpdateQuestion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "question updated",
		"data":    nil,
	})
}

func DeleteQuestion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "question deleted",
		"data":    nil,
	})
}

func GetAllOptions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "options fetch",
		"data":    nil,
	})
}

func GetOption(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "option fetched",
		"data":    nil,
	})
}

func AddOption(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "option added",
		"data":    nil,
	})
}

func UpdateOption(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "option updated",
		"data":    nil,
	})
}

func DeleteOption(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "option deleted",
		"data":    nil,
	})
}
