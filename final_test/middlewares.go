package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func EmailAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userEmail := r.Header.Get("Authentication")
		if userEmail == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		ctx := r.Context()
		r.WithContext(context.WithValue(ctx, "userEmail", userEmail))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func withAnswer(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var a Answer
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		rWithAnswer := r.WithContext(context.WithValue(ctx, "answer", a))
		next(w, rWithAnswer)
	}
}

func withQuestion(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var q Question
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		rWithQuestion := r.WithContext(context.WithValue(ctx, "question", q))
		next(w, rWithQuestion)
	}
}
