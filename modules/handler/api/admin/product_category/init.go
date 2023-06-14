package productcategory

import pcc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product_category"

type ProductCategoryHandler struct {
	productCategoryUsecase pcc.ProductCategoryUsecase
}

func New(productCategoryUsecase pcc.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryUsecase,
	}
}
