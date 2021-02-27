package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/questions", getAllQuestionsHandler).Methods("GET")
	r.HandleFunc("/questions", addQuestionHandler).Methods("POST")
	r.HandleFunc("/questions/{id}", getQuestionHandler).Methods("GET")
	r.HandleFunc("/questions/{id}", updateQuestionHandler).Methods("PUT")
	r.HandleFunc("/questions/{id}", removeQuestionHandler).Methods("DELETE")
	r.HandleFunc("/questions/{questionID}/answer", addAnswerHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
