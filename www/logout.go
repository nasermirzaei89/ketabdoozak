package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

func (h *Handler) logoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := h.cookieStore.Get(r, sessionName)
		if err != nil {
			err = errors.Wrap(err, "failed to get www session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		delete(session.Values, "access_token")
		delete(session.Values, "profile")

		err = session.Save(r, w)
		if err != nil {
			err = errors.Wrap(err, "failed to save session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		logoutURL, err := url.Parse("https://" + h.auth0Domain + "/v2/logout")
		if err != nil {
			err = errors.Wrap(err, "failed to parse logout url")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		returnToURL, err := url.Parse(scheme + "://" + r.Host + h.BaseURL())
		if err != nil {
			err = errors.Wrap(err, "failed to parse return to url")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		params := url.Values{}
		params.Add("returnTo", returnToURL.String())
		params.Add("client_id", h.auth0ClientID)

		logoutURL.RawQuery = params.Encode()

		http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
	}
}
