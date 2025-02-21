package api

import (
	"github.com/NYTimes/gziphandler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

type Handler struct {
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

type HandlerOption func(*Handler)

func RegisterEndpoint(pattern string, handler http.Handler) HandlerOption {
	return func(h *Handler) {
		h.mux.Handle(pattern, handler)
	}
}

func RegisterMiddleware(middleware func(http.Handler) http.Handler) HandlerOption {
	return func(h *Handler) {
		h.middlewares = append(h.middlewares, middleware)
	}
}

func RedirectRootTo(url string) HandlerOption {
	return func(h *Handler) {
		h.mux.Handle("/", http.RedirectHandler(url, http.StatusTemporaryRedirect))
	}
}

func NewHandler(options ...HandlerOption) (*Handler, error) {
	handler := &Handler{
		mux:         http.NewServeMux(),
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}

	handler.middlewares = append(handler.middlewares, gziphandler.GzipHandler)

	handler.middlewares = append(handler.middlewares, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			otelhttp.WithRouteTag(r.Pattern, next).ServeHTTP(w, r)
		})
	})

	for _, option := range options {
		option(handler)
	}

	return handler, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = h.mux
	for _, mw := range h.middlewares {
		handler = mw(handler)
	}

	handler.ServeHTTP(w, r)
}

var _ http.Handler = (*Handler)(nil)
