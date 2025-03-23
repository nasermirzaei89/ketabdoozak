package www

import (
	"github.com/a-h/templ"
	"github.com/gorilla/csrf"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) editItemPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			http.Redirect(w, r, h.BaseURL()+"login", http.StatusTemporaryRedirect)

			return
		}

		itemID := r.PathValue("itemId")

		item, err := h.listingSvc.GetItem(r.Context(), itemID)
		if err != nil {
			if errors.As(err, &listing.ItemWithIDNotFoundError{}) {
				h.notFoundPageHandler()(w, r)

				return
			}

			err = errors.Wrapf(err, "failed to get item with id '%s'", itemID)

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		listLocationsRes, err := h.listingSvc.ListLocations(r.Context())
		if err != nil {
			err = errors.Wrap(err, "failed to list locations")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		head := templates.Head{
			Title: "ویرایش کتاب",
			Meta:  nil,
		}

		templ.Handler(templates.HTML(templates.EditItemPage(item, listLocationsRes.Items, csrf.TemplateField(r)), head)).ServeHTTP(w, r)
	}
}
