package ports

import (
	"context"

	"go.mod/internal/core/domain"
)

type ItemRepository interface {
	GetItemById(ctx context.Context, itemToFind string) error
	GetItems(ctx context.Context, itemToFind string) error
	SaveItem(ctx context.Context, a *domain.Item) error
}
