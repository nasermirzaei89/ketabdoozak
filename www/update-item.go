package www

import (
	"github.com/a-h/templ"
	"github.com/gorilla/csrf"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"strings"
)

func (h *Handler) updateItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		itemID := r.PathValue("itemId")

		err := r.ParseForm()
		if err != nil {
			err := errors.Wrap(err, "failed to parse form")

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		req := populateUpdateItemRequestFromPostForm(r)

		err = h.listingSvc.UpdateItem(r.Context(), itemID, req)
		if err != nil {
			err = errors.Wrapf(err, "failed to update item with id '%s'", itemID)

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

		res, err := h.listingSvc.ListLocations(r.Context())
		if err != nil {
			err = errors.Wrap(err, "failed to list locations")

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		msg := "کتاب با موفقیت ذخیره شد."

		if !req.AsDraft {
			msg = "کتاب برای انتشار ارسال شد."
		}

		templ.Handler(templ.Join(
			templates.EditItemPage(item, res.Items, csrf.TemplateField(r)),
			templates.SuccessMessage(msg),
		)).ServeHTTP(w, r)
	}
}

func populateUpdateItemRequestFromPostForm(r *http.Request) *listing.UpdateItemRequest {
	req := listing.UpdateItemRequest{
		Title:       strings.TrimSpace(r.PostFormValue("title")),
		OwnerName:   strings.TrimSpace(r.PostFormValue("ownerName")),
		LocationID:  strings.TrimSpace(r.PostFormValue("locationId")),
		Types:       nil,
		ContactInfo: nil,
		//#nosec G203 -- Service will sanitize user input
		Description:  template.HTML(strings.TrimSpace(r.PostFormValue("description"))),
		Lent:         false,
		ThumbnailURL: strings.TrimSpace(r.PostFormValue("thumbnailUrl")),
		AsDraft:      r.URL.Query().Has("as-draft") && r.URL.Query().Get("as-draft") != "false",
	}

	if r.PostForm.Has("types") {
		req.Types = make([]listing.ItemType, len(r.PostForm["types"]))
		for i, v := range r.PostForm["types"] {
			req.Types[i] = listing.ItemType(v)
		}
	}

	if r.PostForm.Has("contactInfoType") {
		req.ContactInfo = make([]listing.ItemContactInfo, len(r.PostForm["contactInfoType"]))

		for i := range r.PostForm["contactInfoType"] {
			req.ContactInfo[i] = listing.ItemContactInfo{
				Type:  listing.ItemContactInfoType(r.PostForm["contactInfoType"][i]),
				Value: r.PostForm["contactInfoValue"][i],
			}
		}
	}

	return &req
}
