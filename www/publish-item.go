package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) publishItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		itemID := r.PathValue("itemId")

		err := h.listingSvc.PublishItem(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to publish item with id '%s'", itemID)

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		item, err := h.listingSvc.GetItem(r.Context(), itemID)
		if err != nil {
			if errors.As(err, &listing.ItemWithIDNotFoundError{}) {
				h.notFoundPageHandler()(w, r)

				return
			}

			err = errors.Wrapf(err, "failed to get item with id '%s'", itemID)

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		templ.Handler(templates.SingleItemPage(item)).ServeHTTP(w, r)
	}
}
