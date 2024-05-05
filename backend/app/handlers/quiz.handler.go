package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
	"gorm.io/gorm"
)

type Quiz struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsPublic    bool   `json:"is_public" gorm:"default:false"`
}

func GetAllQuizzes(c *fiber.Ctx) error {
	var quizzes []models.Quiz

	if err := database.DB.Db.Where("is_public = ?", true).Find(&quizzes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quizzes list", nil, quizzes))
}

func GetAllMyQuizzes(c *fiber.Ctx) error {
	owner := c.Locals("user").(utils.SignedDetails)

	var quizzes []models.Quiz
	if err := database.DB.Db.Where("user_id = ?", owner.ID).Find(&quizzes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}
	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "user quizzes", nil, quizzes))
}

func GetQuiz(c *fiber.Ctx) error {
	quiz_id := c.Params("quiz_id")
	var quiz models.Quiz

	if err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz fetched", nil, quiz))
}

func CreateQuiz(c *fiber.Ctx) error {
	var quiz models.Quiz
	owner := c.Locals("user").(utils.SignedDetails)

	if err := c.BodyParser(&quiz); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	quiz.ID = uuid.New()
	fmt.Println(owner.ID)
	id, err := uuid.Parse(owner.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error uuid", err, nil))
	}
	quiz.UserID = id

	if err := database.DB.Db.Create(&quiz).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "quiz created", nil, quiz))
}

func UpdateQuiz(c *fiber.Ctx) error {
	owner := c.Locals("user").(utils.SignedDetails)

	var quizReq Quiz
	if err := c.BodyParser(&quizReq); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	quiz_id := c.Params("quiz_id")
	var quiz models.Quiz

	err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "quiz not found", err, nil))
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	} else if quiz.UserID.String() != owner.ID {
		c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "Unauthorised Request", err, nil))
	}
	quiz.IsPublic = quizReq.IsPublic
	quiz.Category = quizReq.Category
	quiz.Description = quizReq.Description
	quiz.Title = quizReq.Title
	database.DB.Db.Save(&quiz)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz updated", nil, quiz))

}

func DeleteQuiz(c *fiber.Ctx) error {
	owner := c.Locals("user").(utils.SignedDetails)
	quiz_id := c.Params("quiz_id")
	var quiz models.Quiz

	err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "quiz not found", err, nil))
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	} else if owner.ID != quiz.UserID.String() {
		c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "Unauthorised Request", err, nil))
	}
	database.DB.Db.Delete(&quiz)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz deleted", nil, nil))
}
