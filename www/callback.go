package www

import (
	"context"
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
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

		sessionState := session.Values[keyState]
		queryState := r.URL.Query().Get("state")

		if sessionState != queryState {
			err = errors.Errorf("invalid state parameter: session state is '%s', query state is '%s'", sessionState, queryState)

			w.WriteHeader(http.StatusBadRequest)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		delete(session.Values, keyState)

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

		username, ok := profile["preferred_username"].(string)
		if !ok {
			err = errors.New("preferred_username not found in session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		err = h.addUserToAuthenticatedGroup(r.Context(), username)
		if err != nil {
			err = errors.Wrap(err, "failed to add user to authenticated group")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		userSession := &Session{
			ID:           profile["jti"].(string),
			AccessToken:  token.AccessToken,
			TokenType:    token.Type(),
			RefreshToken: token.RefreshToken,
			ExpiresIn:    token.ExpiresIn,
			Expiry:       token.Expiry,
		}

		err = h.sessionRepo.Insert(r.Context(), userSession)
		if err != nil {
			err = errors.Wrap(err, "failed to store user session")

			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(templates.HTML(templates.ErrorPage(err), templates.ErrorHead())).ServeHTTP(w, r)

			return
		}

		session.Values[keySessionID] = userSession.ID

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

func (h *Handler) addUserToAuthenticatedGroup(ctx context.Context, username string) error {
	err := h.authzSvc.AddToGroup(ctx, username, sharedcontext.Authenticated)
	if err != nil {
		return errors.Wrap(err, "failed to add user to authenticated group")
	}

	return nil
}
