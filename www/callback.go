package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) callbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := h.cookieStore.Get(r, sessionName)
		if err != nil {
			err = errors.Wrap(err, "failed to get www session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		sessionState := session.Values["state"]
		queryState := r.URL.Query().Get("state")

		if sessionState != queryState {
			err = errors.New("invalid state parameter")

			w.WriteHeader(http.StatusBadRequest)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		token, err := h.auth.Config.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			err = errors.Wrap(err, "failed to exchange token")

			w.WriteHeader(http.StatusUnauthorized)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		idToken, err := h.auth.VerifyIDToken(r.Context(), token)
		if err != nil {
			err = errors.Wrap(err, "failed to verify ID token")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		var profile map[string]any

		err = idToken.Claims(&profile)
		if err != nil {
			err = errors.Wrap(err, "failed to extract claims")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		session.Values["profile"] = profile

		err = session.Save(r, w)
		if err != nil {
			err = errors.Wrap(err, "failed to save session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		http.Redirect(w, r, h.BaseURL(), http.StatusTemporaryRedirect)
	}
}
