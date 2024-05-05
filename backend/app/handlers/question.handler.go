package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
	"gorm.io/gorm"
)

func GetAllQuestion(c *fiber.Ctx) error {
	quizId := c.Params("quiz_id")
	var questions []models.Question

	if err := database.DB.Db.Where("quiz_id = ?", quizId).Find(&questions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quizzes list", nil, questions))
}

func GetQuestion(c *fiber.Ctx) error {
	question_id := c.Params("question_id")
	var question models.Question

	if err := database.DB.Db.Where("id = ?", question_id).First(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz fetched", nil, question))
}

func CreateQuestion(c *fiber.Ctx) error {
	quiz_id := c.Params("quiz_id")

	quiz := new(models.Quiz)
	owner := c.Locals("user").(utils.SignedDetails)

	if err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "quiz not found", err, nil))
	} else if quiz.UserID.String() != owner.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "not quiz owner", errors.New("unauthorized request"), nil))
	}

	var question models.Question
	if err := c.BodyParser(&question); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "wrong question format", err, nil))
	}

	question.ID = uuid.New()
	question.QuizID = uuid.MustParse(quiz_id)
	if err := database.DB.Db.Create(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	data := fiber.Map{
		"question": question,
	}
	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "question added", nil, data))
}

func UpdateQuestion(c *fiber.Ctx) error {
	var questionReq models.Question
	if err := c.BodyParser(&questionReq); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	owner := c.Locals("user").(utils.SignedDetails)
	quiz_id := c.Params("quiz_id")
	question_id := c.Params("question_id")
	var question models.Question

	// verify is he owner of quiz with quiz_id
	quiz := new(models.Quiz)
	if err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "quiz not found", err, nil))
	} else if quiz.UserID.String() != owner.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "you are not quiz owner", errors.New("unauthorized request"), nil))
	}

	err := database.DB.Db.Where("id = ?", question_id).First(&question).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "no question found with this id", err, nil))
	}

	question.Name = questionReq.Name
	database.DB.Db.Save(&question)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "quiz updated", nil, question))
}

func DeleteQuestion(c *fiber.Ctx) error {
	owner := c.Locals("user").(utils.SignedDetails)
	quiz_id := c.Params("quiz_id")
	question_id := c.Params("question_id")
	var question models.Question

	quiz := new(models.Quiz)
	if err := database.DB.Db.Where("id = ?", quiz_id).First(&quiz).Error; err != nil {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "quiz not found", err, nil))
	} else if quiz.UserID.String() != owner.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "you are not quiz owner", errors.New("unauthorized request"), nil))
	}

	err := database.DB.Db.Where("id = ?", question_id).First(&question).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "question not found", err, nil))
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	database.DB.Db.Delete(&question)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "question deleted", nil, nil))
}

func GetAllOptions(c *fiber.Ctx) error {
	question_id := c.Params("question_id")
	var options []models.Option

	if err := database.DB.Db.Where("question_id = ?", question_id).Find(&options).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "options founds", nil, options))
}

func GetOption(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "option fetched",
		"data":    nil,
	})
}

func CreateOption(c *fiber.Ctx) error {
	question_id, err := uuid.Parse(c.Params("question_id"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "question_id is not correct", err, nil))
	}

	var optionReq models.Option
	if err := c.BodyParser(&optionReq); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "quiz_id is not correct", err, nil))
	}
	option := models.Option{
		ID:         uuid.New(),
		Name:       optionReq.Name,
		QuestionID: question_id,
	}

	if err := database.DB.Db.Create(&option).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "option add", nil, option))
}

func UpdateOption(c *fiber.Ctx) error {
	var optionReq models.Option
	if err := c.BodyParser(&optionReq); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "quiz_id is not correct", err, nil))
	}
	option := models.Option{
		ID:         uuid.New(),
		Name:       optionReq.Name,
		QuestionID: optionReq.QuestionID,
	}

	if err := database.DB.Db.Create(&option).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "option updated", nil, option))
}

func DeleteOption(c *fiber.Ctx) error {
	option_id := c.Params("option_id")

	var option models.Option

	err := database.DB.Db.Where("id = ?", option_id).First(&option).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNoContent).JSON(utils.APIResponse("error", "question not found", err, nil))
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "option deleted", nil, option))
}
