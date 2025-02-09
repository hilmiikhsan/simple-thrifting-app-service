package product

import (
	"context"

	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/interfaces"
	"github.com/hilmiikhsan/thrifting-app-service/internal/models"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	ProductRepo interfaces.IProductRepository
	Logger      *logrus.Logger
}

func (s *ProductService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error {
	err := s.ProductRepo.InsertNewProduct(ctx, models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	})
	if err != nil {
		s.Logger.Error("service::CreateProduct - Failed to insert new product : ", err)
		return err
	}

	return nil
}

func (s *ProductService) GetAllProduct(ctx context.Context) ([]dto.GetProductResponse, error) {
	res := []dto.GetProductResponse{}

	products, err := s.ProductRepo.FindAllProduct(ctx)
	if err != nil {
		s.Logger.Error("service::GetAllProduct - Failed to get all product : ", err)
		return res, err
	}

	for _, product := range products {
		res = append(res, dto.GetProductResponse{
			ID:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		})
	}

	return res, nil
}

func (s *ProductService) GetDetailProduct(ctx context.Context, id string) (*dto.GetProductResponse, error) {
	product, err := s.ProductRepo.FindProductByID(ctx, id)
	if err != nil {
		s.Logger.Error("service::GetDetailProduct - Failed to get product by id : ", err)
		return nil, err
	}

	return &dto.GetProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}
