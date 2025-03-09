package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (h *Handler) newItemPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			http.Redirect(w, r, h.BaseURL()+"login", http.StatusTemporaryRedirect)

			return
		}

		userName, err := h.userFullName(r)
		if err != nil {
			err = errors.Wrap(err, "getting user name")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		item := &listing.Item{
			ID:            "",
			Title:         "",
			OwnerID:       "",
			OwnerName:     userName,
			LocationID:    "",
			LocationTitle: "",
			Types:         nil,
			ContactInfo:   nil,
			Description:   "",
			Status:        "",
			Lent:          false,
			ThumbnailURL:  "",
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			PublishedAt:   nil,
		}

		res, err := h.listingSvc.ListLocations(r.Context())
		if err != nil {
			err = errors.Wrap(err, "failed to list locations")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		head := templates.Head{
			Title: "افزودن کتاب",
			Meta:  nil,
		}

		templ.Handler(templates.HTML(templates.NewItemPage(item, res.Items), head)).ServeHTTP(w, r)
	}
}
