package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) logoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := h.cookieStore.Get(r, sessionName)
		if err != nil {
			err = errors.Wrap(err, "failed to get www session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		delete(session.Values, keySessionID)

		err = session.Save(r, w)
		if err != nil {
			err = errors.Wrap(err, "failed to save session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		switch r.Method {
		case http.MethodGet:
			http.Redirect(w, r, h.oidcLogoutURL, http.StatusTemporaryRedirect)
		case http.MethodPost:
			respond.Done(w, r, nil)
		}
	}
}
