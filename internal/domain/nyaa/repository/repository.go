package repository

import (
	"context"

	"github.com/rl404/nyaa-x-discord/internal/domain/nyaa/entity"
)

// Repository contains functions for nyaa domain.
type Repository interface {
	GetFeeds(ctx context.Context, filter, category string, queries []string) ([]entity.Feed, error)
}
