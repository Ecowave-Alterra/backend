package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() (*[]ep.Product, error)
}

type productRepo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) ProductRepository {
	return &productRepo{
		DB,
	}
}
