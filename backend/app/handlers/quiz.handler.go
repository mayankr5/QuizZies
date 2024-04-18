package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
)

func GetAllQuizzes(c *fiber.Ctx) error {
	var quizzes []models.Quiz

	if err := database.DB.Db.Where("isPublic = ?", true).Find(&quizzes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quizzes list", nil, quizzes))
}

func GetAllMyQuizzes(c *fiber.Ctx) error {
	owner := c.Locals("auth_token").(models.AuthToken)

	var quizzes []models.Quiz
	if err := database.DB.Db.Where("user_id = ?", owner.UserID).Find(&quizzes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}
	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "user quizzes", nil, quizzes))
}

func GetQuiz(c *fiber.Ctx) error {
	cred := c.Params("id")
	var quiz models.Quiz

	if err := database.DB.Db.Where("id = ?", cred).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz fetched", nil, quiz))
}

func CreateQuiz(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "quiz is created",
		"data":    nil,
	})
}

func UpdateQuiz(c *fiber.Ctx) error {
	return nil
}

func DeleteQuiz(c *fiber.Ctx) error {
	return nil
}
