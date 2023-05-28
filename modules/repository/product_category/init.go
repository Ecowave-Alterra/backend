package productcategory

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"gorm.io/gorm"
)

type ProductCategoryRepo interface {
	CreateProductCategory(productCategory *pct.ProductCategory) error
	UpdateProductCategory(productCategory *pct.ProductCategory, id int) error
	DeleteProductCategory(productCategory *pct.ProductCategory, id int) error
	GetAllProductCategory(productCategory *[]pct.ProductCategory) error
	SearchingProductCategoryByName(productCategory *[]pct.ProductCategory, name string) (bool, error)
	IsProductCategoryAvailable(productCategory *pct.ProductCategory, name string) (bool, error)
}

type productCategoryRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProductCategoryRepo {
	return &productCategoryRepo{
		db,
	}
}
