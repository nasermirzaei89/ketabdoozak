package www

import "net/http"

func (h *Handler) registerRoutes() {
	staticServer := http.FileServerFS(static)
	h.mux.Handle("GET /static/", staticServer)

	h.mux.Handle("GET /login", h.loginHandler())
	h.mux.Handle("GET /logout", h.logoutHandler())
	h.mux.Handle("GET /callback", h.callbackHandler())

	h.mux.Handle("GET /", h.indexPageHandler())

	h.mux.Handle("GET /items/{itemId}", h.singleItemPageHandler())
	h.mux.Handle("GET /items/{itemId}/contact-info", h.singleItemGetContactInfoHandler())

	h.mux.Handle("GET /my/items", h.userItemsPageHandler())
	h.mux.Handle("GET /items/new", h.newItemPageHandler())
	//h.mux.Handle("POST /items", h.createItemHandler())
	h.mux.Handle("GET /items/{itemId}/edit", h.editItemPageHandler())
	//h.mux.Handle("PUT /items/{itemId}", h.updateItemHandler())
	h.mux.Handle("DELETE /items/{itemId}", h.deleteItemHandler())

	h.mux.Handle("POST /upload-item-thumbnail", h.uploadItemThumbnailHandler())

	h.mux.Handle("GET /about-us", h.aboutUsPageHandler())
}
