package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/nasermirzaei89/ketabdoozak/www"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type SessionRepo struct {
	rdb *redis.Client
}

var _ www.SessionRepository = (*SessionRepo)(nil)

func NewSessionRepo(rdb *redis.Client) *SessionRepo {
	return &SessionRepo{rdb: rdb}
}

func (repo *SessionRepo) makeKey(id string) string {
	return "www:session:" + id
}

func (repo *SessionRepo) Get(ctx context.Context, id string) (*www.Session, error) {
	res, err := repo.rdb.Get(ctx, repo.makeKey(id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, www.SessionNotFoundError{ID: id}
		}

		return nil, errors.Wrap(err, "failed to get session")
	}

	buf := bytes.NewBufferString(res)

	var session www.Session

	err = gob.NewDecoder(buf).Decode(&session)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode session")
	}

	return &session, nil
}

func (repo *SessionRepo) set(ctx context.Context, id string, session *www.Session) error {
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(session)
	if err != nil {
		return errors.Wrap(err, "failed to encode session")
	}

	err = repo.rdb.Set(ctx, repo.makeKey(id), buf.String(), time.Until(session.Expiry)).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set session")
	}

	return nil
}

func (repo *SessionRepo) Insert(ctx context.Context, session *www.Session) error {
	_, err := repo.Get(ctx, session.ID)
	if err != nil {
		if !errors.As(err, &www.SessionNotFoundError{}) {
			return errors.Wrap(err, "failed to check session in redis")
		}
	} else {
		// TODO: don't use inline error
		return errors.New("session already exists")
	}

	err = repo.set(ctx, session.ID, session)
	if err != nil {
		return errors.Wrap(err, "failed to set session")
	}

	return nil
}

func (repo *SessionRepo) Replace(ctx context.Context, id string, session *www.Session) error {
	err := repo.set(ctx, id, session)
	if err != nil {
		return errors.Wrap(err, "failed to set session")
	}

	return nil
}
