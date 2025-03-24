package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/nasermirzaei89/ketabdoozak/www/templates/utils"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"strings"
)

func (h *Handler) createItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		err := r.ParseForm()
		if err != nil {
			err := errors.Wrap(err, "failed to parse form")

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		req := populateCreateItemRequestFromPostForm(r)

		item, err := h.listingSvc.CreateItem(r.Context(), req)
		if err != nil {
			err = errors.Wrap(err, "failed to create item")

			sendErrorMessage(w, r, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Header().Set("HX-Redirect", string(utils.GetEditItemURL(r.Context(), item.ID)))

		templ.Handler(templates.SuccessMessage("Item has been created successfully.")).ServeHTTP(w, r)
	}
}

func populateCreateItemRequestFromPostForm(r *http.Request) *listing.CreateItemRequest {
	req := listing.CreateItemRequest{
		Title:       strings.TrimSpace(r.PostFormValue("title")),
		OwnerName:   strings.TrimSpace(r.PostFormValue("ownerName")),
		LocationID:  strings.TrimSpace(r.PostFormValue("locationId")),
		Types:       nil,
		ContactInfo: nil,
		//#nosec G203 -- Service will sanitize user input
		Description:  template.HTML(strings.TrimSpace(r.PostFormValue("description"))),
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
