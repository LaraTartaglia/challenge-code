package handlers

import "go.mod/internal/core/domain"

type Item struct {
	Code        string   `json:"code"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	Photos      []string `json:"photos"`
}

func (i Item) toDomain() domain.Item {
	itemToSave := domain.Item{
		Code:        i.Code,
		Title:       i.Title,
		Description: i.Description,
		Price:       i.Price,
		Stock:       i.Stock,
		Photos:      i.Photos,
	}

	return itemToSave
}
