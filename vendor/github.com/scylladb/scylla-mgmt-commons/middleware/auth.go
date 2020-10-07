// Copyright (C) 2017 ScyllaDB

package middleware

import (
	"net/http"

	"github.com/scylladb/scylla-mgmt-commons/httpx"
)

// AddToken sets authorization header. If token is empty it immediately returns
// the next handler.
func AddToken(next http.RoundTripper, token string) http.RoundTripper {
	if token == "" {
		return next
	}

	return httpx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		r := httpx.CloneRequest(req)
		r.Header.Set("Authorization", "Bearer "+token)
		return next.RoundTrip(r)
	})
}
