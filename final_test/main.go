package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func EmailAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userEmail := r.Header.Get("Authentication")
		if userEmail == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func main() {
	r := mux.NewRouter()

	publicQuestionsRoutes := r.PathPrefix("/questions").Methods("GET").Subrouter()
	publicQuestionsRoutes.HandleFunc("", getAllQuestionsHandler).Methods("GET")
	publicQuestionsRoutes.HandleFunc("/{questionId}", getQuestionHandler).Methods("GET")

	authedQuestionsRoutes := r.PathPrefix("/questions").Methods("POST", "PUT", "DELETE").Subrouter()
	authedQuestionsRoutes.Use(EmailAuthMiddleware)
	authedQuestionsRoutes.HandleFunc("", addQuestionHandler).Methods("POST")
	authedQuestionsRoutes.HandleFunc("/{questionId}", updateQuestionHandler).Methods("PUT")
	authedQuestionsRoutes.HandleFunc("/{questionId}", removeQuestionHandler).Methods("DELETE")
	authedQuestionsRoutes.HandleFunc("/{questionId}/answer", addAnswerHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
