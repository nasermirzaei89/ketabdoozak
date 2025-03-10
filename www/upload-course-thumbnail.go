package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
)

func (h *Handler) uploadItemThumbnailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !h.isAuthenticated(r) {
			w.WriteHeader(http.StatusForbidden)
			templ.Handler(templates.ErrorMessage("Access forbidden")).ServeHTTP(w, r)

			return
		}

		filemanager.UploadFileHandler(h.fileManagerSvc).ServeHTTP(w, r)
	}
}
