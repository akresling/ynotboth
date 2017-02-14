package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/akresling/ynotboth/pb"
)

// Router will define the routes for this example service
func Router(s pb.ExampleServer) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloEndpoint(s))
	return mux
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok {
		// Not a Go kit transport error, but a business-logic error.
		// Allows us to encode specific HTTP response codes based on response
		return encodeError(ctx, e, w)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) error {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// codeFrom is used to map an error variable to an http status code
func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
