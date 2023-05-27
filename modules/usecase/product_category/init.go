package productcategory

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
	pcr "github.com/berrylradianh/ecowave-go/modules/repository/product_category"
)

type ProductCategoryUsecase interface {
	CreateProductCategory(productCategory *pct.ProductCategory) error
	UpdateProductCategory(productCategory *pct.ProductCategory, id int) error
	DeleteProductCategory(productCategory *pct.ProductCategory, id int) error
	GetAllProductCategory(productCategory *[]pct.ProductCategory) error
	SearchingProductCategoyByName(productCategory *pct.ProductCategory, name string) error
}

type productCategoryUsecase struct {
	productCategoryRepo pcr.ProductCategoryRepo
}

func New(productCategoryRepo pcr.ProductCategoryRepo) *productCategoryUsecase {
	return &productCategoryUsecase{
		productCategoryRepo,
	}
}
