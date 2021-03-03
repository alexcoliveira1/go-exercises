package answer

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a RESTAPI server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	publicQuestionsRoutes := r.PathPrefix("/questions").Methods("GET").Subrouter()
	publicQuestionsRoutes.HandleFunc("", getAllQuestionsHandler).Methods("GET")
	publicQuestionsRoutes.HandleFunc("/{questionId}", getQuestionHandler).Methods("GET")

	authedQuestionsRoutes := r.PathPrefix("/questions").Methods("POST", "PUT", "DELETE").Subrouter()
	authedQuestionsRoutes.Use(EmailAuthMiddleware)
	authedQuestionsRoutes.HandleFunc("", withQuestion(addQuestionHandler)).Methods("POST")
	authedQuestionsRoutes.HandleFunc("/{questionId}", withQuestion(updateQuestionHandler)).Methods("PUT")
	authedQuestionsRoutes.HandleFunc("/{questionId}", removeQuestionHandler).Methods("DELETE")
	authedQuestionsRoutes.HandleFunc("/{questionId}/answer", withAnswer(addAnswerHandler)).Methods("POST")

	grpcRoutes := r.PathPrefix("/private").Subrouter()
	grpcRoutes.Methods("GET").Path("/gethello").Handler(httptransport.NewServer(
		endpoints.GetHelloEndpoint,
		decodeGetHelloRequest,
		encodeResponse,
	))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", responseHeaderJSON)
		next.ServeHTTP(w, r)
	})
}
