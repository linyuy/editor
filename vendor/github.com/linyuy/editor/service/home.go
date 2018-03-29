package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/static/editor", http.StatusFound)
	}
}
