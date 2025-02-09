package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=5,max=50"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Stock       int     `json:"stock" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10,max=255"`
}

type GetProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}
