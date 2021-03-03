package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	publicQuestionsRoutes := r.PathPrefix("/questions").Methods("GET").Subrouter()
	publicQuestionsRoutes.HandleFunc("", getAllQuestionsHandler).Methods("GET")
	publicQuestionsRoutes.HandleFunc("/{questionId}", getQuestionHandler).Methods("GET")

	authedQuestionsRoutes := r.PathPrefix("/questions").Methods("POST", "PUT", "DELETE").Subrouter()
	authedQuestionsRoutes.Use(EmailAuthMiddleware)
	authedQuestionsRoutes.HandleFunc("", withQuestion(addQuestionHandler)).Methods("POST")
	authedQuestionsRoutes.HandleFunc("/{questionId}", withQuestion(updateQuestionHandler)).Methods("PUT")
	authedQuestionsRoutes.HandleFunc("/{questionId}", removeQuestionHandler).Methods("DELETE")
	authedQuestionsRoutes.HandleFunc("/{questionId}/answer", withAnswer(addAnswerHandler)).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
