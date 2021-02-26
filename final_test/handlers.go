package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	responseHeaderJSON = "application/json; charset=utf-8"
)

// Get one question by its ID
func getQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	q, err := getQuestionByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", responseHeaderJSON)
	if err := json.NewEncoder(w).Encode(q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Get a list of all questions
func getAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userEmail := query.Get("user")

	w.Header().Set("Content-Type", responseHeaderJSON)

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

func getQuestionFromRequestBody(r *http.Request) (*Question, error) {
	var q Question
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		return nil, err
	}
	return &q, nil
}

// Create a new question
func addQuestionHandler(w http.ResponseWriter, r *http.Request) {
	q, err := getQuestionFromRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addQuestion(*q)

	w.Header().Set("Content-Type", responseHeaderJSON)
	w.Write([]byte("{\"success\":true}"))
}

// Get all the questions created by a given user
func getUserQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["user"]

	w.Header().Set("Content-Type", responseHeaderJSON)
	if err := json.NewEncoder(w).Encode(getUserQuestions(user)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Update an existing question (the statement and/or the answer)
// Delete an existing question
