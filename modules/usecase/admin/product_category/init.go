package product_category

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
	pcr "github.com/berrylradianh/ecowave-go/modules/repository/admin/product_category"
)

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory *pct.ProductCategory) (bool, error)
	UpdateProductCategory(productCategory *pct.ProductCategory, id int) (bool, error)
	DeleteProductCategory(productCategory *pct.ProductCategory, id int) error
	GetAllProductCategory(offset, pageSize int) (*[]pct.ProductCategory, int64, error)
	SearchingProductCategoryByName(productCategory *[]pct.ProductCategory, name string) (bool, error)
}

type productCategoryUsecase struct {
	productCategoryRepo pcr.ProductCategoryRepo
}

func New(productCategoryRepo pcr.ProductCategoryRepo) *productCategoryUsecase {
	return &productCategoryUsecase{
		productCategoryRepo,
	}
}
