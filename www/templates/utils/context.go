package utils

import (
	"context"
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
)

type ContextKey string

const ContextKeyBaseURL ContextKey = "baseUrl"

func GetBaseURL(ctx context.Context) templ.SafeURL {
	return templ.SafeURL(ctx.Value(ContextKeyBaseURL).(string))
}

const ContextKeyIsAuthenticated ContextKey = "isAuthenticated"

func IsAuthenticated(ctx context.Context) bool {
	return ctx.Value(ContextKeyIsAuthenticated).(bool)
}

const ContextKeyEnv ContextKey = "env"

func IsProduction(ctx context.Context) bool {
	return ctx.Value(ContextKeyEnv).(string) == "production"
}

func IsCurrentUser(ctx context.Context, username string) bool {
	return sharedcontext.GetSubject(ctx) == username
}
