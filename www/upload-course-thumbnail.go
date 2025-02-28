package www

import (
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"net/http"
)

func (h *Handler) uploadItemThumbnailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filemanager.UploadFileHandler(h.fileManagerSvc).ServeHTTP(w, r)
	}
}
