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

		item, err := h.listingSvc.GetItem(r.Context(), itemID)
		if err != nil {
			err = errors.Wrapf(err, "failed to get item with id '%s'", itemID)

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.ErrorMessage(err.Error())).ServeHTTP(w, r)

			return
		}

		templ.Handler(templates.ContactInfoModal(item.ContactInfo)).ServeHTTP(w, r)
	}
}
