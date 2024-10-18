package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fasttrack/quiz-app/service"
)

// APIResponse is a simple structure for the API gateway layer response.
type APIResponse struct {
	Score      int    `json:"score"`
	Comparison string `json:"comparison"`
}

type Handler struct {
	service service.QuizService
}

// NewHandler creates a new handler with the provided service.
func NewHandler(service service.QuizService) *Handler {
	return &Handler{service: service}
}

// GetQuestions handles the request for fetching questions.
func (h *Handler) GetQuestions(c *gin.Context) {
	ctx := c.Request.Context()

	questions, err := h.service.GetQuestions(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

// SubmitAnswers handles the request for submitting answers and returns the score and comparison.
func (h *Handler) SubmitAnswers(c *gin.Context) {
	ctx := c.Request.Context()

	var userAnswers []int
	if err := c.ShouldBindJSON(&userAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service layer to get the business logic response
	serviceResponse, err := h.service.SubmitAnswers(ctx, userAnswers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert service response to API response
	apiResponse := APIResponse{
		Score:      serviceResponse.Score,
		Comparison: serviceResponse.Comparison,
	}

	c.JSON(http.StatusOK, apiResponse)
}

// AddQuestion handles the request to add a new question.
func (h *Handler) AddQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	var newQuestion service.Question // use the service layer's question structure
	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.service.AddQuestion(ctx, newQuestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Question added successfully"})
}