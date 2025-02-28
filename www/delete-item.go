package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) deleteItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		err := h.listingSvc.DeleteItem(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to delete item with id '%s'", itemID)

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		w.Header().Set("HX-Redirect", h.BaseURL()+"my/items")

		respond.Done(w, r, nil)
	}
}
