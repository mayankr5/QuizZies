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
	user.Get("/leaderboard", middlewares.Authentication, handlers.UserLeaderBoard)

	// Quizzies routes
	quiz := api.Group("/quizzes")
	quiz.Get("/", middlewares.Authentication, handlers.GetAllQuizzes)
	quiz.Get("/my-quizzes", middlewares.Authentication, handlers.GetAllMyQuizzes)
	quiz.Get("/:quiz_id", middlewares.Authentication, handlers.GetQuiz)
	quiz.Post("/", middlewares.Authentication, handlers.CreateQuiz)
	quiz.Put("/:quiz_id", middlewares.Authentication, handlers.UpdateQuiz)
	quiz.Delete("/:quiz_id", middlewares.Authentication, handlers.DeleteQuiz)
	quiz.Post("/:quiz_id/submit", middlewares.Authentication, handlers.SubmitQuiz)

	// Questions routes
	question := quiz.Group("/:quiz_id/questions")
	question.Get("/", middlewares.Authentication, handlers.GetAllQuestion)
	question.Get("/:question_id", middlewares.Authentication, handlers.GetQuestion)
	question.Post("/", middlewares.Authentication, handlers.CreateQuestion)
	question.Put("/:question_id", middlewares.Authentication, handlers.UpdateQuestion)
	question.Delete("/:question_id", middlewares.Authentication, handlers.DeleteQuestion)

	// Options routes
	option := question.Group("/:question_id")
	option.Get("/", middlewares.Authentication, handlers.GetAllOptions)
	option.Get("/:option_id", middlewares.Authentication, handlers.GetOption)
	option.Post("/", middlewares.Authentication, handlers.CreateOption)
	option.Put("/:option_id", middlewares.Authentication, handlers.UpdateOption)
	option.Delete("/:option_id", middlewares.Authentication, handlers.DeleteOption)

	// leaderboard
	api.Get("/leaderboard/:quiz_id", middlewares.Authentication, handlers.LeaderBoard)

}
