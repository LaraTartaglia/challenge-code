package ports

import (
	"context"

	"go.mod/internal/core/domain"
)

type ItemService interface {
	GetItems(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) error
	PostItem(ctx context.Context, a *domain.Item) error
}
