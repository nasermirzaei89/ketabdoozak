package www

import (
	"context"
	"embed"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates/utils"
	"net/http"
)

//go:embed static
var static embed.FS

const serviceName = "www"

const sessionName = "www-session"

type Handler struct {
	mux            *http.ServeMux
	middlewares    []func(http.Handler) http.Handler
	baseURL        templ.SafeURL
	cookieStore    *sessions.CookieStore
	auth           *Authenticator
	auth0Domain    string
	auth0ClientID  string
	fileManagerSvc *filemanager.Service
	listingSvc     *listing.Service
}

func NewHandler(
	baseURL string,
	env string,
	cookieStore *sessions.CookieStore,
	auth *Authenticator,
	auth0Domain string,
	auth0ClientID string,
	fileManagerSvc *filemanager.Service,
	listingSvc *listing.Service,
) (*Handler, error) {
	h := &Handler{
		mux:            http.NewServeMux(),
		middlewares:    make([]func(http.Handler) http.Handler, 0),
		baseURL:        templ.SafeURL(baseURL),
		cookieStore:    cookieStore,
		auth:           auth,
		auth0Domain:    auth0Domain,
		auth0ClientID:  auth0ClientID,
		fileManagerSvc: fileManagerSvc,
		listingSvc:     listingSvc,
	}

	h.middlewares = append(h.middlewares,
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyBaseURL, baseURL))
				next.ServeHTTP(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyIsAuthenticated, h.isAuthenticated(r)))
				next.ServeHTTP(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyEnv, env))
				next.ServeHTTP(w, r)
			})
		},
	)

	h.registerRoutes()

	return h, nil
}

func (h *Handler) BaseURL() string {
	return string(h.baseURL)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = h.mux
	for _, mw := range h.middlewares {
		handler = mw(handler)
	}

	handler.ServeHTTP(w, r)
}

var _ http.Handler = (*Handler)(nil)
