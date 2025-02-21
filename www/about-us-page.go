package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func (h *Handler) aboutUsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.HTML(templates.AboutUsPage())).ServeHTTP(w, r)
	}
}
