package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) uploadItemThumbnailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r, err := h.setRequestContextSubject(r)
		if err != nil {
			err = errors.Wrap(err, "failed to set request context subject")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		filemanager.UploadFileHandler(h.fileManagerSvc).ServeHTTP(w, r)
	}
}
