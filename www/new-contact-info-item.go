package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func (h *Handler) newContactInfoItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		templ.Handler(templates.ContactInfoFormItem("", "")).ServeHTTP(w, r)
	}
}
