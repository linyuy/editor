package service

import (
	"net/http"
)

// NotImplemented replies to the request with an HTTP 501 not implemented error.
func NotImplemented(w http.ResponseWriter, r *http.Request) { http.Error(w, "501 page not implemented", http.StatusNotImplemented) }

// NotImplementedHandler returns a simple request handler
// that replies to each request with a ``501 page not implemented'' reply.
func NotImplementedHandler() http.Handler { return http.HandlerFunc(NotImplemented) }
