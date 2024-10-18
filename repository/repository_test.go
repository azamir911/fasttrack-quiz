package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepository_AddQuestion(t *testing.T) {
	repo := NewRepository()

	question := Question{
		ID:            1,
		QuestionText:  "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}

	// Add the question
	err := repo.AddQuestion(context.Background(), question)
	assert.NoError(t, err, "Error should be nil when adding a valid question")

	// Try adding the same question again (should return an error)
	err = repo.AddQuestion(context.Background(), question)
	assert.EqualError(t, err, ErrQuestionExists.Error(), "Error should be ErrQuestionExists when adding a duplicate question")
}

func TestInMemoryRepository_GetAllQuestions(t *testing.T) {
	repo := NewRepository()

	question1 := Question{
		ID:            2,
		QuestionText:  "What is the largest ocean?",
		Alternatives:  []string{"Atlantic", "Pacific", "Indian", "Arctic"},
		CorrectAnswer: 1,
	}
	question2 := Question{
		ID:            1,
		QuestionText:  "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}

	// Add the questions
	repo.AddQuestion(context.Background(), question1)
	repo.AddQuestion(context.Background(), question2)

	// Get all questions
	questions, err := repo.GetAllQuestions(context.Background())
	assert.NoError(t, err, "Error should be nil when retrieving all questions")

	// Verify that the questions are returned in order by ID
	assert.Len(t, questions, 2, "There should be 2 questions")
	assert.Equal(t, 1, questions[0].ID, "First question should have ID 1")
	assert.Equal(t, 2, questions[1].ID, "Second question should have ID 2")
}

func TestInMemoryRepository_GetQuestionByID(t *testing.T) {
	repo := NewRepository()

	question := Question{
		ID:            1,
		QuestionText:  "What is the capital of France?",
		Alternatives:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		CorrectAnswer: 2,
	}

	// Add the question
	repo.AddQuestion(context.Background(), question)

	// Get the question by ID
	foundQuestion, err := repo.GetQuestionByID(context.Background(), 1)
	assert.NoError(t, err, "Error should be nil when retrieving a valid question by ID")
	assert.Equal(t, question, foundQuestion, "The returned question should match the one added")

	// Try retrieving a non-existing question
	_, err = repo.GetQuestionByID(context.Background(), 999)
	assert.EqualError(t, err, ErrQuestionNotFound.Error(), "Error should be ErrQuestionNotFound for a non-existing question")
}

func TestInMemoryRepository_AddScore(t *testing.T) {
	repo := NewRepository()

	// Add scores
	err := repo.AddScore(context.Background(), 5)
	assert.NoError(t, err, "Error should be nil when adding a score")

	err = repo.AddScore(context.Background(), 8)
	assert.NoError(t, err, "Error should be nil when adding another score")
}
