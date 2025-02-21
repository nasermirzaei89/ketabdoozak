package swagger

import (
	"github.com/swaggo/http-swagger/v2"
	"net/http"
)

type Handler struct {
	mux *http.ServeMux
}

var _ http.Handler = (*Handler)(nil)

func NewHandler() *Handler {
	handler := &Handler{
		mux: http.NewServeMux(),
	}

	handler.RegisterRoutes()

	return handler
}

func (h *Handler) RegisterRoutes() {
	h.mux.Handle("/", httpSwagger.Handler(
		httpSwagger.PersistAuthorization(true),
		httpSwagger.Layout(httpSwagger.BaseLayout),
	))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
