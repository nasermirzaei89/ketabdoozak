package www

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

func (h *Handler) registerRoutes() {
	staticServer := http.FileServerFS(static)
	h.mux.Handle("GET /static/", staticServer)

	h.mux.Handle("GET /auth/login", h.loginHandler())
	h.mux.Handle("GET /auth/logout", h.logoutHandler())
	h.mux.Handle("POST /auth/logout", h.logoutHandler())
	h.mux.Handle("GET /auth/callback", h.callbackHandler())

	h.mux.Handle("GET /", h.indexPageHandler())

	h.mux.Handle("GET /items/{itemId}", h.singleItemPageHandler())
	h.mux.Handle("GET /items/{itemId}/contact-info", h.singleItemGetContactInfoHandler())

	h.mux.Handle("GET /my/items", h.userItemsPageHandler())
	h.mux.Handle("GET /items/new", h.newItemPageHandler())
	h.mux.Handle("POST /items", h.createItemHandler())
	h.mux.Handle("GET /items/{itemId}/edit", h.editItemPageHandler())
	h.mux.Handle("PUT /items/{itemId}", h.updateItemHandler())
	h.mux.Handle("POST /items/{itemId}/send-for-publish", h.sendItemForPublishHandler())
	h.mux.Handle("POST /items/{itemId}/publish", h.publishItemHandler())
	h.mux.Handle("POST /items/{itemId}/archive", h.archiveItemHandler())
	h.mux.Handle("DELETE /items/{itemId}", h.deleteItemHandler())

	h.mux.Handle("POST /upload-item-thumbnail", h.uploadItemThumbnailHandler())

	h.mux.Handle("GET /new-contact-info-item", h.newContactInfoItemHandler())

	h.mux.Handle("GET /about-us", h.aboutUsPageHandler())
}
