package main

import (
	"github.com/google/uuid"
)

type User struct {
	Email     string     `json:"email"`
	Questions []Question `json:"questions"`
}

type Answer struct {
	ID            string   `json:"id"`
	Body          string   `json:"body"`
	UserEmail     string   `json:"userEmail"`
	Votes         int      `json:"votes"`
	UpvotesList   []string `json:"-"`
	DownvotesList []string `json:"-"`
}

type Question struct {
	ID            string   `json:"id"`
	Content       string   `json:"content"`
	Answer        *Answer  `json:"answer"`
	UserEmail     string   `json:"userEmail"`
	Votes         int      `json:"votes"`
	UpvotesList   []string `json:"-"`
	DownvotesList []string `json:"-"`
}

var questions []Question = make([]Question, 0)

type QuestionNotFoundError struct{}

func (m QuestionNotFoundError) Error() string {
	return "Question not found"
}

// Get one question by its ID
func getQuestionByID(id string) (*Question, error) {
	for _, q := range questions {
		if q.ID == id {
			return &q, nil
		}
	}
	return nil, QuestionNotFoundError{}
}

// Get a list of all questions
func getQuestions() []Question {
	return questions
}

// Get all the questions created by a given user
func getUserQuestions(email string) []Question {
	userQuestions := make([]Question, 0)
	for _, q := range questions {
		if q.UserEmail == email {
			userQuestions = append(userQuestions, q)
		}
	}
	return userQuestions
}

type QuestionIdNotEmptyError struct{}

func (m QuestionIdNotEmptyError) Error() string {
	return "ID field must be empty"
}

type NewQuestionWithAnswerError struct{}

func (m NewQuestionWithAnswerError) Error() string {
	return "A new Question must not be created with an answer"
}

func fillQuestionOptionalFields(newQuestion *Question) {
	newQuestion.Votes = 0
	newQuestion.DownvotesList = make([]string, 0)
	newQuestion.UpvotesList = make([]string, 0)
}

func fillNewQuestionWithOldFields(newQuestion *Question, oldQuestion *Question) {
	newQuestion.Votes = oldQuestion.Votes
	newQuestion.DownvotesList = oldQuestion.DownvotesList
	newQuestion.UpvotesList = oldQuestion.UpvotesList
}

// Create a new question
func addQuestion(newQuestion Question) (*Question, error) {
	if newQuestion.ID != "" {
		return nil, QuestionIdNotEmptyError{}
	}
	if newQuestion.Answer != nil {
		return nil, NewQuestionWithAnswerError{}
	}
	newQuestion.ID = uuid.New().String()
	fillQuestionOptionalFields(&newQuestion)
	questions = append(questions, newQuestion)
	return &newQuestion, nil
}

type QuestionIdEmptyError struct{}

func (m QuestionIdEmptyError) Error() string {
	return "ID field cant be empty"
}

// Update an existing question (the statement and/or the answer)
func updateQuestion(updatedQuestion Question) (*Question, error) {
	if updatedQuestion.ID == "" {
		return nil, QuestionIdEmptyError{}
	}

	foundIndex := -1
	for i, q := range questions {
		if q.ID == updatedQuestion.ID {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return nil, QuestionNotFoundError{}
	}

	fillNewQuestionWithOldFields(&updatedQuestion, &questions[foundIndex])

	questions[foundIndex] = updatedQuestion

	return &questions[foundIndex], nil
}

// Delete an existing question
func deleteQuestion(id string) (bool, error) {
	if id == "" {
		return false, QuestionIdEmptyError{}
	}

	foundIndex := -1
	for i, q := range questions {
		if q.ID == id {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return false, QuestionNotFoundError{}
	}

	questions = append(questions[:foundIndex], questions[foundIndex+1:]...)

	return true, nil
}
