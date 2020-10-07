// Copyright (C) 2017 ScyllaDB

package middleware

import (
	"net/http"

	"github.com/scylladb/scylla-mgmt-commons/httpx"
)

// FixScyllaContentType adjusts Scylla REST API response so that it can be consumed
// by Open API.
func FixScyllaContentType(next http.RoundTripper) http.RoundTripper {
	return httpx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		defer func() {
			if resp != nil {
				// Force JSON, Scylla returns "text/plain" that misleads the
				// unmarshaller and breaks processing.
				resp.Header.Set("Content-Type", "application/json")
			}
		}()
		return next.RoundTrip(req)
	})
}
