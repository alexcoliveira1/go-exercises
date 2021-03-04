package answer

import (
	"context"
)

// Service provides some "date capabilities" to your application
type Service interface {
	GetQuestions(ctx context.Context) ([]Question, error)
	GetQuestionByID(ctx context.Context) (Question, error)
	AddQuestion(ctx context.Context) (Question, error)
	UpdateQuestion(ctx context.Context) (Question, error)
	RemoveQuestion(ctx context.Context) (bool, error)
	AddAnswer(ctx context.Context) (Question, error)
}

type answerService struct{}

// NewService makes a new Service.
func NewService() Service {
	return answerService{}
}

// Get will return today's date
func (answerService) GetQuestions(ctx context.Context) ([]Question, error) {
	userEmail, ok := ctx.Value(keyUserEmail).(string)
	if ok && userEmail != "" {
		return getUserQuestions(userEmail), nil
	}
	return getQuestions(), nil
}

// Get will return today's date
func (answerService) GetQuestionByID(ctx context.Context) (Question, error) {
	questionID := ctx.Value(keyQuestionID).(string)
	return getQuestionByID(questionID)
}

// Get will return today's date
func (answerService) AddQuestion(ctx context.Context) (Question, error) {
	newQuestion := ctx.Value(keyNewQuestion).(Question)
	return addQuestion(newQuestion)
}

// Get will return today's date
func (answerService) UpdateQuestion(ctx context.Context) (Question, error) {
	updatedQuestion := ctx.Value(keyUpdatedQuestion).(Question)
	return updateQuestion(updatedQuestion)
}

// Get will return today's date
func (answerService) RemoveQuestion(ctx context.Context) (bool, error) {
	questionID := ctx.Value(keyQuestionID).(string)
	return removeQuestion(questionID)
}

// Get will return today's date
func (answerService) AddAnswer(ctx context.Context) (Question, error) {
	questionID := ctx.Value(keyQuestionID).(string)
	newAnswer := ctx.Value(keyNewAnswer).(Answer)
	return answerQuestion(questionID, newAnswer)
}
