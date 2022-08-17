package repository

import (
	"context"

	"github.com/rl404/nyaa-x-discord/internal/domain/user/entity"
)

// Repository contains functions for user domain.
type Repository interface {
	GetByUserID(ctx context.Context, userID string) (*entity.User, error)
	Create(ctx context.Context, userID, channelID string) error
	UpdateFilterByUserID(ctx context.Context, userID string, filter string) error
	UpdateCategoryByUserID(ctx context.Context, userID string, category string) error
	UpdateQueriesByUserID(ctx context.Context, userID string, queries []string) error
	UpdateSubscribeByUserID(ctx context.Context, userID string, subscribe bool) error
	GetSubscribedUsers(ctx context.Context) ([]*entity.User, error)
}
