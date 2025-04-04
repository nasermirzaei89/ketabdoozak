package www

import (
	"github.com/a-h/templ"
	"github.com/gorilla/csrf"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) sendItemForPublishHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		itemID := r.PathValue("itemId")

		err := h.listingSvc.SendItemForPublish(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to send item with id '%s' for publish", itemID)

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

		templ.Handler(templates.SingleItemPage(item, csrf.Token(r))).ServeHTTP(w, r)
	}
}
