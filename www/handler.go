package www

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/nasermirzaei89/ketabdoozak/www/templates/utils"
	"github.com/nasermirzaei89/services/authorization"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

const sessionName = "www-session"

type Handler struct {
	mux            *http.ServeMux
	middlewares    []func(http.Handler) http.Handler
	baseURL        templ.SafeURL
	cookieStore    *sessions.CookieStore
	sessionRepo    SessionRepository
	auth           *Authenticator
	oidcLogoutURL  string
	authzSvc       *authorization.Service
	fileManagerSvc filemanager.Service
	listingSvc     listing.Service
}

func NewHandler(
	baseURL string,
	env string,
	cookieStore *sessions.CookieStore,
	sessionRepo SessionRepository,
	auth *Authenticator,
	oidcLogoutRedirectURL string,
	csrfAuthKey []byte,
	authzSvc *authorization.Service,
	fileManagerSvc filemanager.Service,
	listingSvc listing.Service,
) (*Handler, error) {
	var claims struct {
		EndSessionEndpoint string `json:"end_session_endpoint"` //nolint:tagliatelle
	}

	err := auth.Provider.Claims(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "error getting auth provider claims")
	}

	logoutURL, err := url.Parse(claims.EndSessionEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse logout url")
	}

	params := url.Values{}
	params.Add("post_logout_redirect_uri", oidcLogoutRedirectURL)
	params.Add("client_id", auth.Config.ClientID)

	logoutURL.RawQuery = params.Encode()

	h := &Handler{
		mux:            http.NewServeMux(),
		middlewares:    make([]func(http.Handler) http.Handler, 0),
		baseURL:        templ.SafeURL(baseURL),
		cookieStore:    cookieStore,
		sessionRepo:    sessionRepo,
		auth:           auth,
		oidcLogoutURL:  logoutURL.String(),
		authzSvc:       authzSvc,
		fileManagerSvc: fileManagerSvc,
		listingSvc:     listingSvc,
	}

	h.middlewares = append(h.middlewares,
		csrf.Protect(
			csrfAuthKey,
			csrf.Secure(env != "development"),
			csrf.HttpOnly(true),
			csrf.Path(baseURL),
			csrf.ErrorHandler(h.csrfErrorHandler()),
		),
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyBaseURL, baseURL))
				next.ServeHTTP(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyEnv, env))
				next.ServeHTTP(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), utils.ContextKeyIsAuthenticated, h.isAuthenticated(r)))

				r, err := h.setRequestContextSubject(r)
				if err != nil {
					err := errors.Wrap(err, "failed to set request context subject")

					w.WriteHeader(http.StatusInternalServerError)
					templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

					return
				}

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
