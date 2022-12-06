package shared

import (
	"context"

	"github.com/doge-verse/easy-upgrade-backend/pkg"
)

type key string

const userKey key = "userID"

// WithUser .
func WithUser(ctx context.Context, user *pkg.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser .
func GetUser(ctx context.Context) (*pkg.User, bool) {
	user, ok := ctx.Value(userKey).(*pkg.User)
	return user, ok
}
