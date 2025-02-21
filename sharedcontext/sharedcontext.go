package sharedcontext

import "context"

const (
	// Anonymous is the guest user id.
	Anonymous = "system:anonymous"

	Authenticated   = "system:authenticated"
	Unauthenticated = "system:unauthenticated"

	GroupRoot = "system:group:root"
)

type ContextKeySubject struct{}

func GetSubject(ctx context.Context) string {
	userID, ok := ctx.Value(ContextKeySubject{}).(string)
	if !ok {
		return Anonymous
	}

	return userID
}

func WithSubject(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeySubject{}, userID)
}

func WithServiceSubject(ctx context.Context, serviceName string) context.Context {
	return WithSubject(ctx, "system:service:"+serviceName)
}
