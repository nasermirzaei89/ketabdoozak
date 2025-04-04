package www

import (
	"github.com/a-h/templ"
	"github.com/gorilla/csrf"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) singleItemPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		preview := r.URL.Query().Get("preview") == "true"

		var (
			item *listing.Item
			err  error
		)

		if preview {
			item, err = h.listingSvc.GetItem(r.Context(), itemID)
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
		} else {
			item, err = h.listingSvc.GetPublishedItem(r.Context(), itemID)
			if err != nil {
				if errors.As(err, &listing.ItemWithIDNotFoundError{}) {
					h.notFoundPageHandler()(w, r)

					return
				}

				err = errors.Wrapf(err, "failed to get published item with id '%s'", itemID)

				w.WriteHeader(http.StatusInternalServerError)
				templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

				return
			}
		}

		head := templates.Head{
			Title: item.Title,
			Meta: []templates.Meta{
				{
					Name:    templates.MetaNameOGTitle,
					Content: item.Title,
				},
				{
					Name:    templates.MetaNameOGImage,
					Content: item.ThumbnailURL,
				},
				{
					Name:    templates.MetaNameTwitterTitle,
					Content: item.Title,
				},
				{
					Name:    templates.MetaNameTwitterImage,
					Content: item.ThumbnailURL,
				},
			},
		}

		templ.Handler(templates.HTML(templates.SingleItemPage(item, csrf.Token(r)), head)).ServeHTTP(w, r)
	}
}
