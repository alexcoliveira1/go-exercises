package answer

import (
	"context"
	"encoding/json"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type getHelloRequest struct{}

type getHelloResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeGetHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getHelloRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
