package answer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get one question by its ID
func getQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["questionId"]

	q, err := getQuestionByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Get a list of all questions
func getAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userEmail := query.Get("user")

	if userEmail != "" {
		if err := json.NewEncoder(w).Encode(getUserQuestions(userEmail)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(getQuestions()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return
}

func getQuestionFromRequestBody(r *http.Request) (Question, error) {
	var q Question
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		return q, err
	}
	return q, nil
}

// Create a new question
func addQuestionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := ctx.Value("question").(Question)

	newQuestion, err := addQuestion(q)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(&newQuestion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Get all the questions created by a given user
func getUserQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["user"]

	if err := json.NewEncoder(w).Encode(getUserQuestions(user)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Update an existing question (the statement and/or the answer)
func updateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	newQ := ctx.Value("question").(Question)
	id := mux.Vars(r)["questionId"]

	if newQ.ID != "" && newQ.ID != id {
		http.Error(w, "ID sent in URL does not match the ID on the request body", http.StatusBadRequest)
		return
	}

	newQ.ID = id

	_, err := updateQuestion(newQ)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("{\"success\":true}"))
}

// Delete an existing question
func removeQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["questionId"]

	_, err := deleteQuestion(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("{\"success\":true}"))
}
