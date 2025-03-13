package filemanager

import (
	"fmt"
	"github.com/nasermirzaei89/problem"
	"github.com/nasermirzaei89/respond"
	"github.com/nasermirzaei89/services/authorization"
	"github.com/pkg/errors"
	"net/http"
)

const UploadKey = "file"

type Handler struct {
	mux            *http.ServeMux
	fileManagerSvc Service
}

var _ http.Handler = (*Handler)(nil)

func NewHandler(fileManagerSvc Service) *Handler {
	h := &Handler{
		mux:            http.NewServeMux(),
		fileManagerSvc: fileManagerSvc,
	}

	h.RegisterRoutes()

	return h
}

func (h *Handler) RegisterRoutes() {
	h.mux.Handle("POST /upload", UploadFileHandler(h.fileManagerSvc))
	h.mux.Handle("GET /files/{filename}", DownloadFileHandler(h.fileManagerSvc))
	h.mux.Handle("DELETE /files/{filename}", DeleteFileHandler(h.fileManagerSvc))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// UploadFileHandler
//
//	@Summary	Upload a file
//	@Tags		filemanager
//	@Security	OAuth2Implicit
//	@Accept		multipart/form-data
//	@Param		file	formData	file	true	"File key"
//	@Router		/filemanager/upload [post]
func UploadFileHandler(fileManagerSvc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formFile, header, err := r.FormFile(UploadKey)
		if err != nil {
			respond.Done(w, r, problem.BadRequest(
				fmt.Sprintf("no file provided at key '%s'", UploadKey),
				problem.WithExtension("error", err.Error()),
			))

			return
		}

		defer func() { _ = formFile.Close() }()

		file, err := fileManagerSvc.UploadFile(r.Context(), header.Filename, formFile, header.Header.Get("Content-Type"))
		if err != nil {
			var accessDeniedError authorization.AccessDeniedError

			switch {
			case errors.As(err, &accessDeniedError):
				respond.Done(w, r, problem.Forbidden(err.Error()))
			default:
				respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on upload file")))
			}

			return
		}

		respond.Done(w, r, file)
	})
}

// DownloadFileHandler
//
//	@Summary	Download a file
//	@Tags		filemanager
//	@Security	OAuth2Implicit
//	@Param		filename	path	string	true	"Filename"
//	@Router		/filemanager/files/{filename} [get]
func DownloadFileHandler(fileManagerSvc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := r.PathValue("filename")

		u, err := fileManagerSvc.GetPublicFileURL(r.Context(), filename)
		if err != nil {
			var (
				accessDeniedError           authorization.AccessDeniedError
				fileByFilenameNotFoundError FileByFilenameNotFoundError
			)

			switch {
			case errors.As(err, &accessDeniedError):
				respond.Done(w, r, problem.Forbidden(err.Error()))
			case errors.As(err, &fileByFilenameNotFoundError):
				respond.Done(w, r, problem.NotFound(fmt.Sprintf("file with filename '%s' does not exist", filename)))
			default:
				respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on download file")))
			}

			return
		}

		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	})
}

// DeleteFileHandler
//
//	@Summary	Delete a file
//	@Tags		filemanager
//	@Security	OAuth2Implicit
//	@Param		filename	path	string	true	"Filename"
//	@Router		/filemanager/files/{filename} [delete]
func DeleteFileHandler(fileManagerSvc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := r.PathValue("filename")

		err := fileManagerSvc.DeleteFile(r.Context(), filename)
		if err != nil {
			var (
				accessDeniedError           authorization.AccessDeniedError
				fileByFilenameNotFoundError FileByFilenameNotFoundError
			)

			switch {
			case errors.As(err, &accessDeniedError):
				respond.Done(w, r, problem.Forbidden(err.Error()))
			case errors.As(err, &fileByFilenameNotFoundError):
				respond.Done(w, r, problem.NotFound(fmt.Sprintf("file with filename '%s' does not exist", filename)))
			default:
				respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on download file")))
			}

			return
		}

		respond.Done(w, r, nil)
	})
}
