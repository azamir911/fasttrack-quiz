## Quiz System

This is a simple Quiz System built with Golang that allows users to answer a set of quiz questions through a REST API. The system includes an in-memory data repository, a service layer for business logic, and an API layer using the [Gin](https://github.com/gin-gonic/gin) web framework.

### Table of Contents

- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
- [Running the Tests](#running-the-tests)

---

### Technologies

- **Golang** - Core programming language
- **Gin** - Web framework for handling HTTP requests
- **Cobra** - CLI framework
- **Testify** - Test library for unit tests

### Project Structure

```plaintext
.
├── api-gateway          # Contains the handlers for the REST API endpoints
│   └── handler.go
├── repository           # Contains the in-memory repository for questions and scores
│   ├── repository.go
│   └── repository_test.go
├── service              # Contains the business logic layer
│   ├── service.go
│   └── service_test.go
├── cmd                  # CLI commands using Cobra
│   └── quiz.go
├── main.go              # Entry point for running the server
└── README.md            # Project documentation
```

### Setup Instructions

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/quiz-system.git
   cd quiz-system
   ```

2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Run the Server**:

   To start the server, run the following command:

   ```bash
   go run main.go
   ```

   The server will be running on `http://localhost:8080`.

4. **Run the CLI**:

   Build the CLI binary using:

   ```bash
   go build -o quiz-cli ./cmd/quiz.go
   ```

   Then, use the CLI to interact with the server. For example, to fetch questions:

   ```bash
   ./quiz-cli get-questions
   ```

### API Endpoints

1. **Get All Questions**
   - **Endpoint**: `GET /questions`
   - **Description**: Retrieve all quiz questions.
   - **Response**: JSON array of questions.

2. **Submit Answers**
   - **Endpoint**: `POST /submit`
   - **Description**: Submit answers to the quiz.
   - **Payload**: JSON array of integers, each representing the selected answer.
   - **Response**: JSON object with the score and comparison message.

3. **Add a New Question**
   - **Endpoint**: `POST /add-question`
   - **Description**: Add a new quiz question.
   - **Payload**:
     ```json
     {
       "id": 1,
       "question": "What is the capital of France?",
       "alternatives": ["Berlin", "Madrid", "Paris", "Rome"],
       "correct_answer": 2
     }
     ```
   - **Response**: Success message.

### Running the Tests

Unit tests are located in each package’s respective `_test.go` files.

To run the tests for the entire project:

```bash
go test ./... -v
```

This will run the tests for all packages, including the service and repository layers.

---

### Contributing

Feel free to open issues or submit pull requests. Any contributions are welcome!

---

### License

This project is licensed under the MIT License.

---

This `README.md` should be enough to get you started on GitHub and provide a clear understanding for anyone who wants to use or contribute to the project. Let me know if you'd like to add anything else!
