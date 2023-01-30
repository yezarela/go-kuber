package usecase

import (
	"context"
	"kuba/models"
	_productRepo "kuba/product/repository"

	"github.com/pkg/errors"
)

type Usecase interface {
	GetProductByID(ctx context.Context, id uint) (models.Product, error)
	ListProduct(ctx context.Context, limit int, offset int) ([]models.Product, error)
	CreateProduct(ctx context.Context, data models.Product) (models.Product, error)
}

type usecase struct {
	productRepo _productRepo.Repository
}

func NewUsecase(repo _productRepo.Repository) Usecase {
	return &usecase{
		productRepo: repo,
	}
}

func (m *usecase) GetProductByID(ctx context.Context, id uint) (models.Product, error) {
	op := "product.Usecase.GetProductByID"

	res, err := m.productRepo.GetProduct(ctx, id)
	if err != nil {
		return res, errors.Wrap(err, op)
	}

	return res, nil
}

func (m *usecase) ListProduct(ctx context.Context, limit int, offset int) ([]models.Product, error) {
	op := "product.Usecase.ListProduct"

	res, err := m.productRepo.ListProduct(ctx, limit, offset)
	if err != nil {
		return res, errors.Wrap(err, op)
	}

	return res, nil
}

func (m *usecase) CreateProduct(ctx context.Context, data models.Product) (models.Product, error) {
	op := "product.Usecase.CreateProduct"

	res, err := m.productRepo.CreateProduct(ctx, data)
	if err != nil {
		return res, errors.Wrap(err, op)
	}

	return res, nil
}
