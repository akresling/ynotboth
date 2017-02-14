package service

import (
	"encoding/json"
	"net/http"

	"github.com/akresling/ynotboth/pb"
)

func helloEndpoint(s pb.ExampleServer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var hr pb.HelloRequest
		err := decoder.Decode(&hr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := s.Hello(r.Context(), &hr)
		if err != nil {
			encodeResponse(r.Context(), w, err)
			return
		}
		encodeResponse(r.Context(), w, resp)
	}
}
