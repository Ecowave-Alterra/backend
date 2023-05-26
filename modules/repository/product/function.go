package product

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (productRepo *productRepo) GetAllProducts() (*[]ep.Product, error) {
	var products []ep.Product
	if err := productRepo.DB.Preload("Product_category", "deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}
