package answer

import (
	"context"
	"time"
)

// Service provides some "date capabilities" to your application
type Service interface {
	GetHello(ctx context.Context) (string, error)
}

type dateService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dateService{}
}

// Get will return today's date
func (dateService) GetHello(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}
