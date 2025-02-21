package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func (h *Handler) notFoundPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		templ.Handler(templates.HTML(templates.NotFoundPage())).ServeHTTP(w, r)
	}
}
