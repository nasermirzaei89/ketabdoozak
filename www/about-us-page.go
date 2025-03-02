package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func (h *Handler) aboutUsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		head := templates.Head{
			Title: "درباره ما",
			Meta:  nil,
		}

		templ.Handler(templates.HTML(templates.AboutUsPage(), head)).ServeHTTP(w, r)
	}
}
