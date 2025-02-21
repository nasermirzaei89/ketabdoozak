package utils

import (
	"context"
	"github.com/a-h/templ"
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
