package www

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			err = errors.Wrap(err, "failed to generate random state")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		session, err := h.cookieStore.Get(r, sessionName)
		if err != nil {
			err = errors.Wrap(err, "failed to get www session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		session.Values["state"] = state

		err = session.Save(r, w)
		if err != nil {
			err = errors.Wrap(err, "failed to save session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err))).ServeHTTP(w, r)

			return
		}

		http.Redirect(w, r, h.auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
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
