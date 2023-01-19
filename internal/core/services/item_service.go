package services

import (
	"context"
	"fmt"

	"go.mod/internal/core/domain"
	"go.mod/internal/core/ports"
)

type itemService struct {
	repo ports.ItemRepository
}

func NewItemService(repo ports.ItemRepository) (ports.ItemService, error) {
	if repo == nil {
		return nil, fmt.Errorf("erro with repo")
	}
	return &itemService{
		repo: repo,
	}, nil
}

func (i *itemService) GetItems(ctx context.Context, id string) error {
	return i.repo.GetItems(ctx, id)
}

func (i *itemService) GetById(ctx context.Context, id string) error {
	return i.repo.GetItemById(ctx, id)
}

func (i *itemService) PostItem(ctx context.Context, body *domain.Item) error {
	return i.repo.SaveItem(ctx, body)
}
