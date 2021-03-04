package answer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetQuestionsEndpoint   endpoint.Endpoint
	GetQuestionEndpoint    endpoint.Endpoint
	AddQuestionEndpoint    endpoint.Endpoint
	UpdateQuestionEndpoint endpoint.Endpoint
	RemoveQuestionEndpoint endpoint.Endpoint
	AddAnswerEndpoint      endpoint.Endpoint
}

type BadRequestError struct{}

func (e BadRequestError) Error() string {
	return "Bad request"
}

// MakeGetQuestionsEndpoint returns the response from our service "GetQuestions"
func MakeGetQuestionsEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getQuestionsRequest)

		if !ok {
			err := BadRequestError{}
			return getQuestionsResponse{nil, err.Error()}, err
		}

		userCtx := context.WithValue(ctx, keyUserEmail, req.UserEmail)

		questions, err := srv.GetQuestions(userCtx)
		if err != nil {
			return getQuestionsResponse{questions, err.Error()}, nil
		}
		return getQuestionsResponse{questions, ""}, nil
	}
}

type QuestionIDIsRequiredError struct{}

func (e QuestionIDIsRequiredError) Error() string {
	return "Question id is required"
}

// MakeGetQuestionsEndpoint returns the response from our service "GetQuestionByID"
func MakeGetQuestionEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getQuestionRequest)

		if !ok {
			err := BadRequestError{}
			return getQuestionResponse{nil, err.Error()}, err
		}

		if req.QuestionID == "" {
			err := QuestionIDIsRequiredError{}
			return getQuestionResponse{nil, err.Error()}, err
		}

		idCtx := context.WithValue(ctx, keyQuestionID, req.QuestionID)

		question, err := srv.GetQuestionByID(idCtx)
		if err != nil {
			return getQuestionResponse{nil, err.Error()}, err
		}

		return getQuestionResponse{&question, ""}, nil
	}
}

// MakeAddQuestionEndpoint returns the response from our service "addQuestion"
func MakeAddQuestionEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(addQuestionRequest)

		if !ok {
			err := BadRequestError{}
			return addQuestionResponse{nil, err.Error()}, err
		}

		qCtx := context.WithValue(ctx, keyNewQuestion, req.Question)

		question, err := srv.AddQuestion(qCtx)
		if err != nil {
			return addQuestionResponse{nil, err.Error()}, err
		}

		return addQuestionResponse{&question, ""}, nil
	}
}

// MakeUpdateQuestionEndpoint returns the response from our service "updateQuestion"
func MakeUpdateQuestionEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateQuestionRequest)

		if !ok {
			err := BadRequestError{}
			return updateQuestionResponse{nil, err.Error()}, err
		}

		qCtx := context.WithValue(ctx, keyUpdatedQuestion, req.Question)

		question, err := srv.UpdateQuestion(qCtx)
		if err != nil {
			return updateQuestionResponse{nil, err.Error()}, err
		}

		return updateQuestionResponse{&question, ""}, nil
	}
}

// MakeRemoveQuestionEndpoint returns the response from our service "RemoveQuestion"
func MakeRemoveQuestionEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(removeQuestionRequest)

		if !ok {
			err := BadRequestError{}
			return removeQuestionResponse{false, err.Error()}, err
		}

		idCtx := context.WithValue(ctx, keyQuestionID, req.QuestionID)

		deleted, err := srv.RemoveQuestion(idCtx)
		if err != nil {
			return removeQuestionResponse{false, err.Error()}, err
		}

		return removeQuestionResponse{deleted, ""}, nil
	}
}

// MakeAddAnswerEndpoint returns the response from our service "AddAnswer"
func MakeAddAnswerEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(addAnswerRequest)

		if !ok {
			err := BadRequestError{}
			return addAnswerResponse{nil, err.Error()}, err
		}

		idCtx := context.WithValue(ctx, keyQuestionID, req.QuestionID)
		aCtx := context.WithValue(idCtx, keyNewAnswer, req.Answer)

		answeredQuestion, err := srv.AddAnswer(aCtx)
		if err != nil {
			return addAnswerResponse{nil, err.Error()}, err
		}

		return addAnswerResponse{&answeredQuestion, ""}, nil
	}
}
