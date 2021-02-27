package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getAnswerFromRequestBody(r *http.Request) (Answer, error) {
	var a Answer
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return a, err
	}
	return a, nil
}

// Get one question by its ID
func addAnswerHandler(w http.ResponseWriter, r *http.Request) {
	questionID := mux.Vars(r)["questionID"]
	a, err := getAnswerFromRequestBody(r)

	q, err := answerQuestion(questionID, a)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", responseHeaderJSON)
	if err := json.NewEncoder(w).Encode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
