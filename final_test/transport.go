package answer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type getQuestionsRequest struct {
	UserEmail string
}

type getQuestionsResponse struct {
	Questions []Question `json:"questions"`
	Err       string     `json:"err,omitempty"`
}

type getQuestionRequest struct {
	QuestionID string
}

type getQuestionResponse struct {
	Question *Question `json:"question,omitempty"`
	Err      string    `json:"err,omitempty"`
}

type addQuestionRequest struct {
	Question Question
}

type addQuestionResponse = getQuestionResponse

type updateQuestionRequest = addQuestionRequest

type updateQuestionResponse = getQuestionResponse

type removeQuestionRequest = getQuestionRequest

type removeQuestionResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"err,omitempty"`
}

type addAnswerRequest struct {
	QuestionID string
	Answer     Answer
}

type addAnswerResponse = addQuestionResponse

func decodeGetQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getQuestionsRequest
	query := r.URL.Query()
	userEmail := query.Get("user")
	req.UserEmail = userEmail
	return req, nil
}

func decodeGetQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getQuestionRequest
	id := mux.Vars(r)["questionId"]

	req.QuestionID = id
	return req, nil
}

func decodeAddQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addQuestionRequest

	var q Question
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		return req, err
	}

	req.Question = q
	return req, nil
}

func decodeUpdateQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeAddQuestionRequest(ctx, r)
}

func decodeRemoveQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeGetQuestionRequest(ctx, r)
}

func decodeAddAnswerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addAnswerRequest

	id := mux.Vars(r)["questionId"]

	var a Answer
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return req, err
	}

	req.QuestionID = id
	req.Answer = a
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
