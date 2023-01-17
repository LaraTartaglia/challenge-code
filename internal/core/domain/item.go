package domain

type Item struct {
	ID          int
	Code        string
	Title       string
	Description string
	Price       int
	Stock       int
	Photos      []string
	CreatedAt   string
	UpdatedAt   string
	Status      string
}
