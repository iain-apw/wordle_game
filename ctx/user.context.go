package ctx

import (
	"context"

	"github.com/iain-apw/wordle_game/models"
)

type key int

const (
	userContextID key = iota
)

func AddUserToContext(user *models.User, ctx context.Context) (context.Context, error) {
	updatedContext := context.WithValue(ctx, userContextID, user)
	return updatedContext, nil
}

func GetUserFromContext(ctx context.Context) (*models.User, error) {
	user := ctx.Value(userContextID).(*models.User)
	return user, nil
}
