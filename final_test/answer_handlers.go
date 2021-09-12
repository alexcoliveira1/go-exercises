package answer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get one question by its ID
func addAnswerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	questionId := mux.Vars(r)["questionId"]

	a := ctx.Value("answer").(Answer)

	q, err := answerQuestion(questionId, a)

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
