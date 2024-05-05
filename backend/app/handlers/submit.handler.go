package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
)

func SubmitQuiz(c *fiber.Ctx) error {
	var score models.ScoreBoard
	if err := c.BodyParser(&score); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.APIResponse("error", "request is not correct format", err, nil))
	}

	owner := c.Locals("user").(utils.SignedDetails)
	quiz_id := c.Params("quiz_id")

	score.ID = uuid.New()
	score.QuizID = uuid.Must(uuid.FromBytes([]byte(quiz_id)))
	score.UserID = uuid.Must(uuid.FromBytes([]byte(owner.ID)))

	scoreModel := new(models.ScoreBoard)
	if err := database.DB.Db.Where("user_id = ? and quiz_id = ?", owner.ID, quiz_id).First(&scoreModel).Error; err != nil {
		database.DB.Db.Create(&score)
	} else {
		scoreModel.Score = score.Score
		database.DB.Db.Save(&scoreModel)
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "score updated", nil, nil))
}
