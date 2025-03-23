package www

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

func (h *Handler) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			err = errors.Wrap(err, "failed to generate random state")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		session, err := h.cookieStore.Get(r, sessionName)
		if err != nil {
			err = errors.Wrap(err, "failed to get www session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		session.Values[keyState] = state

		err = session.Save(r, w)
		if err != nil {
			err = errors.Wrap(err, "failed to save session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		http.Redirect(w, r, h.auth.Config.AuthCodeURL(state, oauth2.AccessTypeOffline), http.StatusFound)
	}
}

func generateRandomState() (string, error) {
	const stateLen = 32

	b := make([]byte, stateLen)

	_, err := rand.Read(b)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate random state")
	}

	state := base64.RawStdEncoding.EncodeToString(b)

	return state, nil
}
