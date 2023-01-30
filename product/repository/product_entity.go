package repository

import (
	"kuba/models"

	"gorm.io/gorm"
)

type ProductEntity struct {
	gorm.Model
	Name        string
	Description string
	Email       string
}

func (p ProductEntity) TableName() string {
	return "products"
}

func (p ProductEntity) toProductModel() models.Product {
	return models.Product{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
