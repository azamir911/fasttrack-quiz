package main

import (
	"context"
	"fmt"
	"log"

	"fasttrack/quiz-app/api-gateway"
	"fasttrack/quiz-app/repository"
	"fasttrack/quiz-app/service"
	"github.com/gin-gonic/gin"
)

func main2() {
	// Initialize the repository
	repo := repository.NewRepository()

	// Initialize the service with the repository
	svc := service.NewQuizService(repo)

	// Initialize the handler with the service
	handler := apigateway.NewHandler(svc)

	ctx := context.Background()

	repo.AddQuestion(ctx, repository.Question{ID: 1, QuestionText: "What is the capital of France?", Alternatives: []string{"Berlin", "Madrid", "Paris", "Rome"}, CorrectAnswer: 2})
	repo.AddQuestion(ctx, repository.Question{ID: 2, QuestionText: "What is 2 + 2?", Alternatives: []string{"3", "4", "5", "6"}, CorrectAnswer: 1})
	repo.AddQuestion(ctx, repository.Question{ID: 3, QuestionText: "Which planet is known as the Red Planet?", Alternatives: []string{"Earth", "Venus", "Mars", "Jupiter"}, CorrectAnswer: 2})
	repo.AddQuestion(ctx, repository.Question{ID: 4, QuestionText: "What is the largest ocean on Earth?", Alternatives: []string{"Atlantic", "Indian", "Arctic", "Pacific"}, CorrectAnswer: 3})
	repo.AddQuestion(ctx, repository.Question{ID: 5, QuestionText: "Who wrote 'Hamlet'?", Alternatives: []string{"Mark Twain", "William Shakespeare", "J.K. Rowling", "Ernest Hemingway"}, CorrectAnswer: 1})
	repo.AddQuestion(ctx, repository.Question{ID: 6, QuestionText: "What is the smallest prime number?", Alternatives: []string{"0", "1", "2", "3"}, CorrectAnswer: 2})
	repo.AddQuestion(ctx, repository.Question{ID: 7, QuestionText: "Which gas do plants absorb?", Alternatives: []string{"Oxygen", "Nitrogen", "Carbon Dioxide", "Hydrogen"}, CorrectAnswer: 3})
	repo.AddQuestion(ctx, repository.Question{ID: 8, QuestionText: "What is the hardest natural substance on Earth?", Alternatives: []string{"Gold", "Iron", "Diamond", "Platinum"}, CorrectAnswer: 3})
	repo.AddQuestion(ctx, repository.Question{ID: 9, QuestionText: "What is the chemical symbol for Gold?", Alternatives: []string{"Au", "Ag", "Pb", "Fe"}, CorrectAnswer: 1})
	repo.AddQuestion(ctx, repository.Question{ID: 10, QuestionText: "How many continents are there?", Alternatives: []string{"5", "6", "7", "8"}, CorrectAnswer: 2})

	// Set up the Gin router
	router := gin.Default()

	// Define the routes
	router.GET("/questions", handler.GetQuestions)
	router.POST("/submit", handler.SubmitAnswers)
	router.POST("/add-question", handler.AddQuestion)

	// Start the Gin server
	fmt.Println("Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
