package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) singleItemGetContactInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		item, err := h.listingSvc.GetPublishedItem(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to get published item with id '%s'", itemID)

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		templ.Handler(templates.ContactInfoModal(item.ContactInfo)).ServeHTTP(w, r)
	}
}
