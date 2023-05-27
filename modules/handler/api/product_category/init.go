package productcategory

import pcc "github.com/berrylradianh/ecowave-go/modules/usecase/product_category"

type ProductCategoryHandler struct {
	productCategoryUsecase pcc.ProductCategoryUsecase
}

func New(productCategoryUsecase pcc.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryUsecase,
	}
}
