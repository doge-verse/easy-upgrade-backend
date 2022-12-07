package shared

import (
	"context"

	"github.com/doge-verse/easy-upgrade-backend/models"
)

type key string

const userKey key = "userID"

// WithUser .
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser .
func GetUser(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userKey).(*models.User)
	return user, ok
}
