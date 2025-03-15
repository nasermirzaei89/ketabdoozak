package www

import (
	"context"
	"fmt"
	"time"
)

type Session struct {
	ID           string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
	ExpiresIn    int64
}

type SessionNotFoundError struct {
	ID string
}

func (err SessionNotFoundError) Error() string {
	return fmt.Sprintf("session with id '%s' not found", err.ID)
}

type SessionRepository interface {
	Insert(ctx context.Context, session *Session) (err error)
	Get(ctx context.Context, id string) (session *Session, err error)
	Replace(ctx context.Context, id string, session *Session) (err error)
}
