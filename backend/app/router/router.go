package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/app/handlers"
	"github.com/mayankr5/quizzies/app/middlewares"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Authentication Routes
	user := api.Group("/users")
	user.Post("/register", handlers.Register)
	user.Post("/login", handlers.Login)
	user.Get("/logout", middlewares.Authentication, handlers.Logout)

	//User Routes
	// user.Post("/update", handlers.UpdateUser)
	// user.Delete("/delete", handlers.DeleteUser)
	user.Get("/:user_id", handlers.UserLeaderBoard)

	// Quizzies routes
	quiz := api.Group("/quizzes")
	quiz.Get("/", middlewares.Authentication, handlers.GetAllQuizzes)
	quiz.Get("/my-quizzes", middlewares.Authentication, handlers.GetAllMyQuizzes)
	quiz.Get("/:quiz_id", middlewares.Authentication, handlers.GetQuiz)
	quiz.Post("/", middlewares.Authentication, handlers.CreateQuiz)
	quiz.Put("/:quiz_id", middlewares.Authentication, handlers.UpdateQuiz)
	quiz.Delete("/:quiz_id", middlewares.Authentication, handlers.DeleteQuiz)

	// Questions routes
	question := quiz.Group("/:quiz_id/questions")
	question.Get("/", handlers.GetAllQuestion)
	question.Get("/:question_id", handlers.GetQuestion)
	question.Post("/", handlers.CreateQuestion)
	question.Put("/:question_id", handlers.UpdateQuestion)
	question.Delete("/:question_id", handlers.DeleteQuestion)

	// Options routes
	option := question.Group("/:question_id")
	option.Get("/", handlers.GetAllOptions)
	option.Get("/:option_id", handlers.GetOption)
	option.Post("/", handlers.CreateOption)
	option.Put("/:option_id", handlers.UpdateOption)
	option.Delete("/:option_id", handlers.DeleteOption)

	// leaderboard
	api.Get("/leaderboard/:quiz_id", handlers.LeaderBoard)

}
