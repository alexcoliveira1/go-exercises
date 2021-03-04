package answer

import (
	"context"
	"testing"
)

func TestGetQuestions(t *testing.T) {
	srv, ctx := setup()
	q, err := srv.GetQuestions(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing today's date
	ok := 0 == len(q)
	if !ok {
		t.Errorf("expected questions to be empty")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
