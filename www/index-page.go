package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) indexPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			h.notFoundPageHandler()(w, r)

			return
		}

		q := r.URL.Query().Get("q")

		res, err := h.listingSvc.ListPublishedItems(r.Context(), q)
		if err != nil {
			err = errors.Wrap(err, "failed to list published items")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		templ.Handler(templates.HTML(templates.IndexPage(res.Items, q), templates.EmptyHead())).ServeHTTP(w, r)
	}
}
