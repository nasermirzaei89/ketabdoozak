package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) sendItemForPublishHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			w.WriteHeader(http.StatusForbidden)
			templ.Handler(templates.ErrorMessage("Access forbidden")).ServeHTTP(w, r)

			return
		}

		itemID := r.PathValue("itemId")

		err := h.listingSvc.SendItemForPublish(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to send item with id '%s' for publish", itemID)

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.ErrorMessage(err.Error())).ServeHTTP(w, r)

			return
		}

		item, err := h.listingSvc.GetItem(r.Context(), itemID)
		if err != nil {
			if errors.As(err, &listing.ItemWithIDNotFoundError{}) {
				h.notFoundPageHandler()(w, r)

				return
			}

			err = errors.Wrapf(err, "failed to get item with id '%s'", itemID)

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.ErrorMessage(err.Error())).ServeHTTP(w, r)

			return
		}

		templ.Handler(templates.SingleItemPage(item)).ServeHTTP(w, r)
	}
}
