package www

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
}

// NewAuthenticator instantiates the *Authenticator.
func NewAuthenticator(ctx context.Context, oidcIssuerURL, oidcClientID, oidcClientSecret, oidcRedirectURL string) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		ctx,
		oidcIssuerURL,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create oidc provider")
	}

	conf := oauth2.Config{
		ClientID:     oidcClientID,
		ClientSecret: oidcClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  oidcRedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.Config.ClientID,
	}

	idToken, err := a.Provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify ID token")
	}

	return idToken, nil
}

func (h *Handler) isAuthenticated(r *http.Request) bool {
	session, err := h.cookieStore.Get(r, sessionName)
	if err != nil {
		panic(err)
	}

	return session.Values["profile"] != nil
}

func (h *Handler) userName(r *http.Request) string {
	session, err := h.cookieStore.Get(r, sessionName)
	if err != nil {
		panic(err)
	}

	sessionProfile, ok := session.Values["profile"].(map[string]any)
	if ok {
		userName, ok := sessionProfile["name"].(string)
		if ok {
			return userName
		}
	}

	return ""
}

func (h *Handler) setRequestContextSubject(r *http.Request) (*http.Request, error) {
	session, err := h.cookieStore.Get(r, sessionName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get www session")
	}

	sessionProfile, ok := session.Values["profile"].(map[string]any)
	if ok {
		username, ok := sessionProfile["preferred_username"].(string)
		if ok {
			r = r.WithContext(sharedcontext.WithSubject(r.Context(), username))
		} else {
			r = r.WithContext(sharedcontext.WithSubject(r.Context(), sharedcontext.Anonymous))
		}
	}

	return r, nil
}
