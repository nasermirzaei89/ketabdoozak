package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) userItemsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		ctx := r.Context()

		req := &listing.ListItemsRequest{
			Query:   q,
			OwnerID: sharedcontext.GetSubject(ctx),
		}

		res, err := h.listingSvc.ListItems(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "failed to list items")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		templ.Handler(templates.HTML(templates.UserItemsPage(res.Items, q))).ServeHTTP(w, r)
	}
}
