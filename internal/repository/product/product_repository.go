package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *ProductRepository) InsertNewProduct(ctx context.Context, product models.Product) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryInsertNewProduct),
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)
	if err != nil {
		r.Logger.Error("repo::InsertNewProduct - Failed to insert new product : ", err)
		return err
	}

	return nil
}

func (r *ProductRepository) FindAllProduct(ctx context.Context) ([]models.Product, error) {
	var (
		products []models.Product
	)

	err := r.DB.SelectContext(ctx, &products, r.DB.Rebind(queryFindAllProduct))
	if err != nil {
		r.Logger.Error("repo::GetAllProduct - Failed to get all product : ", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindProductByID(ctx context.Context, id string) (*models.Product, error) {
	var res = new(models.Product)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindProductByID), id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("repo::GetProductByID - Product not found : ", err)
			return nil, errors.New(constants.ErrProductNotFound)
		}

		r.Logger.Error("repo::GetProductByID - Failed to get product by id : ", err)
		return nil, err
	}

	return res, nil
}
