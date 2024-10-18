package service

import (
	"context"
	"fasttrack/quiz-app/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizService_GetQuestions(t *testing.T) {
	// Use the real repository
	repo := repository.NewRepository()
	svc := NewQuizService(repo)

	// Define questions to add to the repository
	question1 := repository.Question{
		ID:            1,
		QuestionText:  "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}
	question2 := repository.Question{
		ID:            2,
		QuestionText:  "Which is the largest ocean?",
		Alternatives:  []string{"Atlantic", "Pacific", "Indian", "Arctic"},
		CorrectAnswer: 1,
	}

	// Add questions to the repository
	err := repo.AddQuestion(context.Background(), question1)
	assert.NoError(t, err)

	err = repo.AddQuestion(context.Background(), question2)
	assert.NoError(t, err)

	// Call the service method
	questions, err := svc.GetQuestions(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, questions, 2, "There should be 2 questions")
	assert.Equal(t, "What is the capital of France?", questions[0].Question, "The first question text should match")
	assert.Equal(t, "Which is the largest ocean?", questions[1].Question, "The second question text should match")
}

func TestQuizService_AddQuestion(t *testing.T) {
	// Use the real repository
	repo := repository.NewRepository()
	svc := NewQuizService(repo)

	// Define the input question
	newQuestion := Question{
		ID:            1,
		Question:      "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}

	// Call the service method
	err := svc.AddQuestion(context.Background(), newQuestion)

	// Assertions
	assert.NoError(t, err, "Adding a valid question should not return an error")

	// Verify that the question was added
	questions, err := repo.GetAllQuestions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, questions, 1, "There should be 1 question in the repository")
	assert.Equal(t, newQuestion.Question, questions[0].QuestionText, "The question text should match")
}

func TestQuizService_SubmitAnswers(t *testing.T) {
	// Use the real repository
	repo := repository.NewRepository()
	svc := NewQuizService(repo)

	// Define a question and add it to the repository
	question := repository.Question{
		ID:            1,
		QuestionText:  "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}
	err := repo.AddQuestion(context.Background(), question)
	assert.NoError(t, err)

	// Call the service method to submit answers
	submitResponse, err := svc.SubmitAnswers(context.Background(), []int{2})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, submitResponse.Score, "The score should be 1")
	assert.Contains(t, submitResponse.Comparison, "You were better than", "The comparison message should be generated")

	// Verify that the score was added to the repository
	scores, err := repo.GetAllScores(context.Background())
	assert.NoError(t, err)
	assert.Len(t, scores, 1, "There should be 1 score")
}

func TestQuizService_SubmitAnswers_Error(t *testing.T) {
	// Use the real repository without adding any questions
	repo := repository.NewRepository()
	svc := NewQuizService(repo)

	// Call the service method (should return an error since no questions exist)
	_, err := svc.SubmitAnswers(context.Background(), []int{2})

	// Assertions
	assert.Error(t, err, "Fetching questions should return an error when no questions exist")
}
