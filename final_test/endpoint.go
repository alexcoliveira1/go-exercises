package answer

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetHelloEndpoint endpoint.Endpoint
}

// MakeGetHelloEndpoint returns the response from our service "getHello"
func MakeGetHelloEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getHelloRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetHello(ctx)
		if err != nil {
			return getHelloResponse{d, err.Error()}, nil
		}
		return getHelloResponse{d, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) GetHello(ctx context.Context) (string, error) {
	req := getHelloRequest{}
	resp, err := e.GetHelloEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getHelloResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}
