package main

import (
	"sync"

	"github.com/google/uuid"
)

var questions []Question = make([]Question, 0)
var questionsLock = sync.RWMutex{}

type QuestionNotFoundError struct{}

func (m QuestionNotFoundError) Error() string {
	return "Question not found"
}

type QuestionIdNotEmptyError struct{}

func (m QuestionIdNotEmptyError) Error() string {
	return "ID field must be empty"
}

type NewQuestionWithAnswerError struct{}

func (m NewQuestionWithAnswerError) Error() string {
	return "A new Question must not be created with an answer"
}

type QuestionIdEmptyError struct{}

func (m QuestionIdEmptyError) Error() string {
	return "ID field cant be empty"
}

type QuestionAlreadyAnsweredError struct{}

func (m QuestionAlreadyAnsweredError) Error() string {
	return "Question already has an answer. To change answer remove the current one and add a new one afterwards. To modify content run an update."
}

// Get one question by its ID
func getQuestionByID(id string) (Question, error) {
	questionsLock.RLock()
	defer questionsLock.RUnlock()

	for _, q := range questions {
		if q.ID == id {
			return q, nil
		}
	}
	return Question{}, QuestionNotFoundError{}
}

// Get a list of all questions
func getQuestions() []Question {
	questionsLock.RLock()
	defer questionsLock.RUnlock()
	return questions
}

// Get all the questions created by a given user
func getUserQuestions(email string) []Question {
	userQuestions := make([]Question, 0)

	questionsLock.RLock()
	defer questionsLock.RUnlock()

	for _, q := range questions {
		if q.UserEmail == email {
			userQuestions = append(userQuestions, q)
		}
	}
	return userQuestions
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
func addQuestion(newQuestion Question) (Question, error) {
	if newQuestion.ID != "" {
		return Question{}, QuestionIdNotEmptyError{}
	}
	if newQuestion.Answer != nil {
		return Question{}, NewQuestionWithAnswerError{}
	}

	newQuestion.ID = uuid.New().String()
	fillQuestionOptionalFields(&newQuestion)

	questionsLock.Lock()
	defer questionsLock.Unlock()

	questions = append(questions, newQuestion)

	return newQuestion, nil
}

// Update an existing question (the statement and/or the answer)
func updateQuestion(updatedQuestion Question) (Question, error) {
	if updatedQuestion.ID == "" {
		return Question{}, QuestionIdEmptyError{}
	}

	questionsLock.Lock()
	defer questionsLock.Unlock()

	foundIndex := -1
	for i, q := range questions {
		if q.ID == updatedQuestion.ID {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return Question{}, QuestionNotFoundError{}
	}

	fillNewQuestionWithOldFields(&updatedQuestion, &questions[foundIndex])

	questions[foundIndex] = updatedQuestion

	return questions[foundIndex], nil
}

// Delete an existing question
func deleteQuestion(id string) (bool, error) {
	if id == "" {
		return false, QuestionIdEmptyError{}
	}

	questionsLock.Lock()
	defer questionsLock.Unlock()

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

func fillAnswerOptionalFields(newAnswer *Answer) {
	newAnswer.Votes = 0
	newAnswer.DownvotesList = make([]string, 0)
	newAnswer.UpvotesList = make([]string, 0)
}

// Answer Question
func answerQuestion(questionID string, answer Answer) (Question, error) {
	q, err := getQuestionByID(questionID)

	if err != nil {
		return Question{}, err
	}

	if q.Answer != nil {
		return Question{}, QuestionAlreadyAnsweredError{}
	}

	fillAnswerOptionalFields(&answer)
	q.Answer = &answer

	return updateQuestion(q)
}
