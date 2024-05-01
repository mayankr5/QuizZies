package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
)

func LeaderBoard(c *fiber.Ctx) error {
	quiz_id := c.Params("quiz_id")
	var scoreboard []models.ScoreBoard

	if err := database.DB.Db.Order("score desc").Where("quiz_id = ?", quiz_id).Find(&scoreboard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "leaderboard of quiz", nil, scoreboard))
}

func UserLeaderBoard(c *fiber.Ctx) error {
	user_id := c.Params("user_id")

	var scoreBoard []models.ScoreBoard

	if err := database.DB.Db.Order("created_at desc").Where("user_id = ?", user_id).Find(&scoreBoard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on database", err, nil))
	}
	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "leaderboard of quiz", nil, scoreBoard))
}
