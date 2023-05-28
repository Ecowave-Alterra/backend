package productcategory

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
	pcr "github.com/berrylradianh/ecowave-go/modules/repository/product_category"
)

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory *pct.ProductCategory) (bool, error)
	UpdateProductCategory(productCategory *pct.ProductCategory, id int) (bool, error)
	DeleteProductCategory(productCategory *pct.ProductCategory, id int) error
	GetAllProductCategory(productCategory *[]pct.ProductCategory) error
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
