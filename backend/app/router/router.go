package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/quizzies/app/handlers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	// Authentication Routes
	user := api.Group("/user")
	user.Post("/register", handlers.Register)
	user.Post("/login", handlers.Login)
	user.Get("/logout", handlers.Logout)

	// Quizzies routes
	quiz := api.Group("/quiz")
	quiz.Get("/", handlers.GetAllQuizzes)
	quiz.Get("/my-quizzes", handlers.GetAllMyQuizzes)
	quiz.Get("/{id}", handlers.GetQuiz)
	quiz.Post("/", handlers.CreateQuiz)
	quiz.Put("/{id}", handlers.UpdateQuiz)
	quiz.Delete("/{id}", handlers.DeleteQuiz)

	// Questions routes
	question := quiz.Group("/{quiz_id}/question")
	question.Get("/", handlers.GetAllQuestion)
	question.Get("/{id}", handlers.GetQuestion)
	question.Post("/", handlers.AddQuestion)
	question.Put("/{id}", handlers.UpdateQuestion)
	question.Delete("/{id}", handlers.DeleteQuestion)

	// Options routes
	option := question.Group("/{question_id}")
	option.Get("/", handlers.GetAllOptions)
	option.Get("/{id}", handlers.GetOption)
	option.Post("/", handlers.AddOption)
	option.Put("/{id}", handlers.UpdateOption)
	option.Delete("/{id}", handlers.DeleteOption)
}
