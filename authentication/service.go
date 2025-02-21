package authentication

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/nasermirzaei89/problem"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"strings"
)

type Service struct {
	verifier *oidc.IDTokenVerifier
	logger   *slog.Logger
	tracer   trace.Tracer
}

func NewService(ctx context.Context, issuer, clientID string) (*Service, error) {
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize oidc provider")
	}

	config := &oidc.Config{
		ClientID:             clientID,
		SupportedSigningAlgs: nil,
		SkipClientIDCheck:    true,
		SkipExpiryCheck:      false,
		SkipIssuerCheck:      false,
		Now:                  nil,
	}

	return &Service{
		verifier: provider.Verifier(config),
		logger:   defaultLogger,
		tracer:   defaultTracer,
	}, nil
}

func (svc *Service) AuthenticateMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := svc.tracer.Start(r.Context(), "AuthenticateMiddleware")
			defer span.End()

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				span.SetAttributes(attribute.String("subject", sharedcontext.Anonymous))

				r = r.WithContext(sharedcontext.WithSubject(ctx, sharedcontext.Anonymous))

				next.ServeHTTP(w, r)

				return
			}

			// TODO: security issue?
			span.SetAttributes(attribute.String("authHeader", authHeader))

			// Extract the token from "Bearer <token>"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				respond.Done(w, r, problem.Unauthorized("invalid authorization header format"))

				return
			}

			token := tokenParts[1]

			// Verify the token
			idToken, err := svc.verifier.Verify(ctx, token)
			if err != nil {
				span.SetStatus(codes.Error, err.Error())

				respond.Done(w, r, problem.Unauthorized("invalid token"))

				return
			}

			// Token is valid, pass to next handler
			r = r.WithContext(sharedcontext.WithSubject(ctx, idToken.Subject))
			next.ServeHTTP(w, r)
		})
	}
}
