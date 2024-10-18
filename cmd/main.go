package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "A CLI to interact with the quiz API",
}

// addQuestionCmd represents the add-question command
var addQuestionCmd = &cobra.Command{
	Use:   "add-question",
	Short: "Add a new question to the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			fmt.Println("Usage: add-question <id> <question> <correct_answer_index> <alternative_1> <alternative_2> ... <alternative_n>")
			os.Exit(1)
		}

		// Parse the arguments
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid question ID:", args[0])
			os.Exit(1)
		}

		questionText := args[1]
		correctAnswerIndex, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid correct answer index:", args[2])
			os.Exit(1)
		}

		alternatives := args[3:]

		// Create the question structure
		question := map[string]interface{}{
			"id":             id,
			"question":       questionText,
			"correct_answer": correctAnswerIndex,
			"alternatives":   alternatives,
		}

		// Convert question to JSON
		questionJSON, err := json.Marshal(question)
		if err != nil {
			fmt.Println("Error encoding question:", err)
			os.Exit(1)
		}

		// Send POST request to the API
		resp, err := http.Post("http://localhost:8080/add-question", "application/json", bytes.NewBuffer(questionJSON))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
			os.Exit(1)
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		if resp.StatusCode == http.StatusCreated {
			fmt.Println("Question added successfully!")
		} else {
			fmt.Println("Failed to add question. Status code:", resp.StatusCode)
		}
	},
}

var getQuestionsCmd = &cobra.Command{
	Use:   "get-questions",
	Short: "Fetches quiz questions",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/questions")
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

var submitAnswersCmd = &cobra.Command{
	Use:   "submit-answers [answers]",
	Short: "Submit your answers (provide them as space-separated integers)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the answers from CLI arguments
		var answers []string
		for _, ans := range args {
			answers = append(answers, ans)
		}

		// Convert the answers to JSON format
		answersStr := "[" + strings.Join(answers, ", ") + "]"

		// Make the POST request
		resp, err := http.Post("http://localhost:8080/submit", "application/json", strings.NewReader(answersStr))
		if err != nil {
			fmt.Println("Error submitting answers:", err)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		// Read and print the response
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(addQuestionCmd)
	rootCmd.AddCommand(getQuestionsCmd)
	rootCmd.AddCommand(submitAnswersCmd)
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
