package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) userItemsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			http.Redirect(w, r, h.BaseURL()+"auth/login", http.StatusTemporaryRedirect)

			return
		}

		ctx := r.Context()

		q := r.URL.Query().Get("q")

		res, err := h.listingSvc.ListMyItems(ctx, q)
		if err != nil {
			err = errors.Wrap(err, "failed to list my items")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		head := templates.Head{
			Title: "کتاب‌های من",
			Meta:  nil,
		}

		templ.Handler(templates.HTML(templates.UserItemsPage(res.Items, q), head)).ServeHTTP(w, r)
	}
}
