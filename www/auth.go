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

const keySessionID = "sessionID"
const keyState = "state"

type NewAuthenticatorRequest struct {
	OIDCIssuerURL    string
	OIDCClientID     string
	OIDCClientSecret string
	OIDCRedirectURL  string
}

// NewAuthenticator instantiates the *Authenticator.
func NewAuthenticator(ctx context.Context, req NewAuthenticatorRequest) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		ctx,
		req.OIDCIssuerURL,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create oidc provider")
	}

	conf := oauth2.Config{
		ClientID:     req.OIDCClientID,
		ClientSecret: req.OIDCClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  req.OIDCRedirectURL,
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
		ClientID:                   a.Config.ClientID,
		SupportedSigningAlgs:       nil,
		SkipClientIDCheck:          false,
		SkipExpiryCheck:            false,
		SkipIssuerCheck:            false,
		Now:                        nil,
		InsecureSkipSignatureCheck: false,
	}

	idToken, err := a.Provider.VerifierContext(ctx, oidcConfig).Verify(ctx, rawIDToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify ID token")
	}

	return idToken, nil
}

func (h *Handler) sessionID(r *http.Request) string {
	session, err := h.cookieStore.Get(r, sessionName)
	if err != nil {
		panic(err)
	}

	sessionID, ok := session.Values[keySessionID].(string)
	if !ok {
		return ""
	}

	return sessionID
}

var ErrNoSessionID = errors.New("no session ID found")

func (h *Handler) session(r *http.Request) (*Session, error) {
	sessionID := h.sessionID(r)
	if sessionID == "" {
		return nil, ErrNoSessionID
	}

	session, err := h.sessionRepo.Get(r.Context(), sessionID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get session")
	}

	return session, nil
}

func (h *Handler) isAuthenticated(r *http.Request) bool {
	_, err := h.session(r)

	return err == nil
}

func (h *Handler) userInfo(r *http.Request) (*oidc.UserInfo, error) {
	session, err := h.session(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get session")
	}

	token := &oauth2.Token{
		AccessToken:  session.AccessToken,
		TokenType:    session.TokenType,
		RefreshToken: session.RefreshToken,
		Expiry:       session.Expiry,
		ExpiresIn:    session.ExpiresIn,
	}

	userInfo, err := h.auth.Provider.UserInfo(r.Context(), h.auth.Config.TokenSource(r.Context(), token))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}

	return userInfo, nil
}

func (h *Handler) userFullName(r *http.Request) (string, error) {
	userInfo, err := h.userInfo(r)
	if err != nil {
		return "", errors.Wrap(err, "failed to get user info")
	}

	var claims struct {
		Name string `json:"name"`
	}

	err = userInfo.Claims(&claims)
	if err != nil {
		return "", errors.Wrap(err, "failed to get user info claims")
	}

	return claims.Name, nil
}

func (h *Handler) username(r *http.Request) (string, error) {
	userInfo, err := h.userInfo(r)
	if err != nil {
		if !errors.Is(err, ErrNoSessionID) && !errors.As(err, &SessionNotFoundError{}) {
			return "", errors.Wrap(err, "failed to get user info")
		}
	} else {
		var claims map[string]any

		err = userInfo.Claims(&claims)
		if err != nil {
			return "", errors.Wrap(err, "failed to get user info claims")
		}

		username, ok := claims["preferred_username"].(string)
		if ok {
			return username, nil
		}
	}

	return sharedcontext.Anonymous, nil
}

func (h *Handler) setRequestContextSubject(r *http.Request) (*http.Request, error) {
	username, err := h.username(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get username")
	}

	r = r.WithContext(sharedcontext.WithSubject(r.Context(), username))

	return r, nil
}
