package repository

import (
	"context"
	"kuba/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository interface {
	GetProduct(ctx context.Context, id uint) (models.Product, error)
	ListProduct(ctx context.Context, limit int, offset int) ([]models.Product, error)
	CreateProduct(ctx context.Context, data models.Product) (models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (m *repository) GetProduct(ctx context.Context, id uint) (res models.Product, err error) {
	op := "product.Repository.GetProduct"

	item := ProductEntity{}
	err = m.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return res, errors.Wrap(err, op)
	}

	return item.toProductModel(), nil
}

func (m *repository) ListProduct(ctx context.Context, limit int, offset int) ([]models.Product, error) {
	op := "product.Repository.ListProduct"

	items := []ProductEntity{}
	err := m.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&items).Error
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	products := []models.Product{}
	for _, item := range items {
		products = append(products, item.toProductModel())
	}

	return products, nil
}

func (m *repository) CreateProduct(ctx context.Context, data models.Product) (res models.Product, err error) {
	op := "product.Repository.CreateProduct"

	item := ProductEntity{
		Name:        data.Name,
		Description: data.Description,
	}

	err = m.db.WithContext(ctx).Create(&item).Error
	if err != nil {
		return res, errors.Wrap(err, op)
	}

	return item.toProductModel(), nil
}
