package www

import (
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) deleteItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		itemID := r.PathValue("itemId")

		err := h.listingSvc.DeleteItem(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to delete item with id '%s'", itemID)

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Header().Set("HX-Redirect", h.BaseURL()+"my/items")

		respond.Done(w, r, nil)
	}
}
