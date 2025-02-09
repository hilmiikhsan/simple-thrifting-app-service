package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/models"
)

type IProductRepository interface {
	InsertNewProduct(ctx context.Context, product models.Product) error
	FindAllProduct(ctx context.Context) ([]models.Product, error)
	FindProductByID(ctx context.Context, id string) (*models.Product, error)
}

type IProductService interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error
	GetAllProduct(ctx context.Context) ([]dto.GetProductResponse, error)
	GetDetailProduct(ctx context.Context, id string) (*dto.GetProductResponse, error)
}

type IProductHandler interface {
	CreateProduct(ctx *gin.Context)
	GetAllProduct(ctx *gin.Context)
	GetDetailProduct(ctx *gin.Context)
}
