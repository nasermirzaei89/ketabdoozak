package www

import (
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"net/http"
)

func (h *Handler) uploadItemThumbnailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			sendErrorMessage(w, r, "Access forbidden", http.StatusForbidden)

			return
		}

		filemanager.UploadFileHandler(h.fileManagerSvc).ServeHTTP(w, r)
	}
}
