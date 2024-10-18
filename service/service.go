package service

import (
	"context"
	"errors"
	"fmt"

	"fasttrack/quiz-app/repository"
)

// SubmitResponse holds the result of the quiz submission in the service layer.
type SubmitResponse struct {
	Score      int
	Comparison string
}

// Question represents the question structure used across the service layer.
type Question struct {
	ID            int      `json:"id"`
	Question      string   `json:"question"`
	Alternatives  []string `json:"alternatives"`
	CorrectAnswer int      `json:"correct_answer"`
}

// QuizService defines the business logic for the quiz.
type QuizService interface {
	GetQuestions(ctx context.Context) ([]Question, error)
	SubmitAnswers(ctx context.Context, answers []int) (SubmitResponse, error)
	AddQuestion(ctx context.Context, question Question) error
}

type QuizServiceImpl struct {
	repo repository.Repository
}

// NewQuizService creates a new instance of QuizService with the given repository.
func NewQuizService(repo repository.Repository) QuizService {
	return &QuizServiceImpl{repo: repo}
}

// GetQuestions fetches all the quiz questions from the repository and maps them to the service layer's question.
func (q *QuizServiceImpl) GetQuestions(ctx context.Context) ([]Question, error) {
	repoQuestions, err := q.repo.GetAllQuestions(ctx)
	if err != nil {
		return nil, err
	}

	serviceQuestions := make([]Question, 0, len(repoQuestions))

	// Map repository questions to service questions
	for _, repoQuestion := range repoQuestions {
		serviceQuestions = append(serviceQuestions, Question{
			ID:            repoQuestion.ID,
			Question:      repoQuestion.QuestionText,
			Alternatives:  repoQuestion.Alternatives,
			CorrectAnswer: repoQuestion.CorrectAnswer,
		})
	}

	return serviceQuestions, nil
}

// SubmitAnswers checks the user's answers and calculates the score.
func (q *QuizServiceImpl) SubmitAnswers(ctx context.Context, answers []int) (SubmitResponse, error) {
	questions, err := q.GetQuestions(ctx)
	if err != nil {
		return SubmitResponse{}, err
	}

	// Return an error if no questions are available
	if len(questions) == 0 {
		return SubmitResponse{}, errors.New("no questions available")
	}

	correctCount := 0

	for i, answer := range answers {
		if i < len(questions) && answer == questions[i].CorrectAnswer {
			correctCount++
		}
	}

	// Calculate comparison against other users
	scores, err := q.repo.GetAllScores(ctx)
	if err != nil {
		return SubmitResponse{}, err
	}

	percentage, err := calculateComparison(scores, correctCount)
	if err != nil {
		return SubmitResponse{}, err
	}

	var comparison string
	if percentage == -1 {
		comparison = "You are the first to do the quiz"
	} else {
		comparison = fmt.Sprintf("You were better than %d%% of all quizzers", percentage)
	}

	// Store the score in the repository
	err = q.repo.AddScore(ctx, correctCount)
	if err != nil {
		return SubmitResponse{}, err
	}

	return SubmitResponse{
		Score:      correctCount,
		Comparison: comparison,
	}, nil
}

// AddQuestion converts the service layer question to the repository format and adds it.
func (q *QuizServiceImpl) AddQuestion(ctx context.Context, question Question) error {
	repoQuestion := repository.Question{
		ID:            question.ID,
		QuestionText:  question.Question,
		Alternatives:  question.Alternatives,
		CorrectAnswer: question.CorrectAnswer,
	}
	return q.repo.AddQuestion(ctx, repoQuestion)
}

// calculateComparison compares the user's score against all other scores.
func calculateComparison(scores []int, userScore int) (int, error) {
	if len(scores) == 0 {
		return -1, nil
	}

	betterCount := 0
	for _, score := range scores {
		if userScore > score {
			betterCount++
		}
	}

	percentage := int(float64(betterCount) / float64(len(scores)) * 100)
	return percentage, nil
}
