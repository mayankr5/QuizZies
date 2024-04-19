package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
	"gorm.io/gorm"
)

func GetAllQuizzes(c *fiber.Ctx) error {
	var quizzes []models.Quiz

	if err := database.DB.Db.Where("is_public = ?", true).Find(&quizzes).Error; err != nil {
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
	quiz_id := c.Params("id")
	var quiz models.Quiz

	if err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz fetched", nil, quiz))
}

func CreateQuiz(c *fiber.Ctx) error {
	var quiz models.Quiz
	if err := c.BodyParser(&quiz); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	data := fiber.Map{}
	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "quiz created", nil, data))
}

func UpdateQuiz(c *fiber.Ctx) error {
	var quizReq models.Quiz
	if err := c.BodyParser(&quizReq); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	quiz_id := c.Params("id")
	var quiz models.Quiz

	err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "no quiz found with this id", err, nil))
	}
	quiz.IsPublic = quizReq.IsPublic
	quiz.Category = quizReq.Category
	quiz.Description = quizReq.Description
	quiz.Title = quizReq.Title
	database.DB.Db.Save(&quiz)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz updated", nil, quizReq))

}

func DeleteQuiz(c *fiber.Ctx) error {
	quiz_id := c.Params("id")
	var quiz models.Quiz

	err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "no quiz found with this id", err, nil))
	}
	database.DB.Db.Delete(&quiz)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz deleted", nil, nil))
}
