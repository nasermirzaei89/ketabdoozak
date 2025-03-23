package www

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func (h *Handler) csrfErrorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendErrorMessage(w, r, csrf.FailureReason(r).Error(), http.StatusForbidden)
	}
}
