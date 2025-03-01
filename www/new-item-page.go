package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) newItemPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			http.Redirect(w, r, h.BaseURL()+"login", http.StatusTemporaryRedirect)

			return
		}

		res, err := h.listingSvc.ListLocations(r.Context())
		if err != nil {
			err = errors.Wrap(err, "failed to list locations")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		templ.Handler(templates.HTML(templates.NewItemPage(res.Items))).ServeHTTP(w, r)
	}
}
