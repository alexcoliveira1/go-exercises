package answer

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a RESTAPI server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {

	getQuestionsHandler := httptransport.NewServer(
		endpoints.GetQuestionsEndpoint,
		decodeGetQuestionsRequest,
		encodeResponse,
	)

	getQuestionHandler := httptransport.NewServer(
		endpoints.GetQuestionEndpoint,
		decodeGetQuestionRequest,
		encodeResponse,
	)

	addQuestionHandler := httptransport.NewServer(
		endpoints.AddQuestionEndpoint,
		decodeAddQuestionRequest,
		encodeResponse,
	)

	updateQuestionHandler := httptransport.NewServer(
		endpoints.UpdateQuestionEndpoint,
		decodeAddQuestionRequest, // reusing from another endpoint
		encodeResponse,
	)

	removeQuestionHandler := httptransport.NewServer(
		endpoints.RemoveQuestionEndpoint,
		decodeGetQuestionRequest, // reusing from another endpoint
		encodeResponse,
	)

	addAnswerHandler := httptransport.NewServer(
		endpoints.AddAnswerEndpoint,
		decodeAddAnswerRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	publicQuestionsRoutes := r.PathPrefix("/questions").Methods("GET").Subrouter()
	publicQuestionsRoutes.Methods("GET").Path("").Handler(getQuestionsHandler)
	publicQuestionsRoutes.Methods("GET").Path("/{questionId}").Handler(getQuestionHandler)

	authedQuestionsRoutes := r.PathPrefix("/questions").Methods("POST", "PUT", "DELETE").Subrouter()
	authedQuestionsRoutes.Use(EmailAuthMiddleware)
	authedQuestionsRoutes.Methods("POST").Path("").Handler(addQuestionHandler)
	authedQuestionsRoutes.Methods("PUT").Path("/{questionId}").Handler(updateQuestionHandler)
	authedQuestionsRoutes.Methods("DELETE").Path("/{questionId}").Handler(removeQuestionHandler)
	authedQuestionsRoutes.Methods("POST").Path("/{questionId}/answer").Handler(addAnswerHandler)

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", responseHeaderJSON)
		next.ServeHTTP(w, r)
	})
}
