package answer

import (
	"context"
	"testing"
	"time"
)

func TestGetHello(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.GetHello(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	timeNow := time.Now()
	today := timeNow.Format("02/01/2006")

	// testing today's date
	ok := today == d
	if !ok {
		t.Errorf("expected dates to be equal")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
