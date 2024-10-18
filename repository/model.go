package repository

// Question represents a question in the repository layer.
type Question struct {
	ID            int
	QuestionText  string
	Alternatives  []string
	CorrectAnswer int
}
