package repository

import (
	"context"
	"errors"
	"sort"
)

// Repository defines the methods to access and modify quiz data.
type Repository interface {
	GetAllQuestions(ctx context.Context) ([]Question, error)
	GetQuestionByID(ctx context.Context, id int) (Question, error)
	AddQuestion(ctx context.Context, question Question) error
	GetAllScores(ctx context.Context) ([]int, error)
	AddScore(ctx context.Context, score int) error
}

var (
	ErrQuestionNotFound = errors.New("question not found")
	ErrQuestionExists   = errors.New("question already exists")
)

type inMemoryRepository struct {
	questions map[int]Question // Map to store questions with question ID as key
	scores    []int            // Slice to store scores
}

// NewRepository creates a new in-memory repository.
func NewRepository() Repository {
	return &inMemoryRepository{
		questions: make(map[int]Question),
		scores:    []int{},
	}
}

// AddQuestion adds a new question to the repository.
func (im *inMemoryRepository) AddQuestion(ctx context.Context, question Question) error {
	// Check if the question already exists by ID
	if _, exists := im.questions[question.ID]; exists {
		return ErrQuestionExists
	}

	// Validate the question data
	if question.QuestionText == "" || len(question.Alternatives) == 0 {
		return errors.New("invalid question: question text and alternatives are required")
	}

	im.questions[question.ID] = question
	return nil
}

// GetAllQuestions returns all quiz questions as a sorted slice.
func (im *inMemoryRepository) GetAllQuestions(ctx context.Context) ([]Question, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// Convert the map to a slice
		questionsSlice := make([]Question, 0, len(im.questions))
		for _, question := range im.questions {
			questionsSlice = append(questionsSlice, question)
		}

		// Sort the slice by question ID
		sort.Slice(questionsSlice, func(i, j int) bool {
			return questionsSlice[i].ID < questionsSlice[j].ID
		})

		return questionsSlice, nil
	}
}

// GetQuestionByID returns a question by its ID.
func (im *inMemoryRepository) GetQuestionByID(ctx context.Context, id int) (Question, error) {
	select {
	case <-ctx.Done():
		return Question{}, ctx.Err()
	default:
		if question, exists := im.questions[id]; exists {
			return question, nil
		}
		return Question{}, ErrQuestionNotFound
	}
}

// AddScore adds a user's score to the repository.
func (im *inMemoryRepository) AddScore(ctx context.Context, score int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		im.scores = append(im.scores, score)
		return nil
	}
}

// GetAllScores returns all the stored quiz scores.
func (im *inMemoryRepository) GetAllScores(ctx context.Context) ([]int, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return im.scores, nil
	}
}
