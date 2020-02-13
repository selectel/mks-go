package testutils

import (
	"fmt"
	"net/http"
	"testing"
)

// HandleReqOpts represents options for the testing utils package handlers.
type HandleReqOpts struct {
	// Mux represents HTTP Mux for a testing handler.
	Mux *http.ServeMux

	// URL represents handler's HTTP URL.
	URL string

	// RawResponse represents raw string HTTP response that needs to be returned
	// by the handler.
	RawResponse string

	// RawRequest represents raw string HTTP request that needs to be compared
	// with the actual request that will be provided by the caller.
	RawRequest string

	// Method contains HTTP method that needs to be compared against real method
	// provided by the caller.
	Method string

	// Status represents HTTP status that will be returned by the handler.
	Status int

	// CallFlag can be used to check if caller sent a request to a handler.
	CallFlag *bool
}

// HandleReqWithoutBody provides the HTTP endpoint to test requests without body.
func HandleReqWithoutBody(t *testing.T, opts *HandleReqOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		fmt.Fprint(w, opts.RawResponse)

		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		*opts.CallFlag = true
	})
}
