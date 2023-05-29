package product

import (
	"fmt"
	"net/http"

	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req ep.ProductRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to bind data",
			"error":   err,
		})
	}
	productDescription := ep.ProductDescription{
		Description: req.Description,
	}
	err = h.productUseCase.CreateProductDescription(&productDescription)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create product description",
			"error":   err,
		})
	}

	status := "tersedia"
	if req.Stock == 0 {
		status = "habis"
	}

	product := ep.Product{
		Product_category_id:    req.Product_category_id,
		Product_description_id: productDescription.ID,
		Name:                   req.Name,
		Stock:                  req.Stock,
		Price:                  req.Price,
		Status:                 status,
		Rating:                 0.00,
	}
	err = h.productUseCase.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create product",
			"error":   err,
		})
	}

	for _, url := range req.Product_image_url {
		productImage := ep.ProductImage{
			Product_id:        product.ID,
			Product_image_url: url,
		}
		err = h.productUseCase.CreateProductImage(&productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to create product image",
				"error":   err,
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully create new product",
	})
}

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProduct(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get product datas",
			"error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Product_Description.Description,
			Product_image_url: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get all product datas",
		"products": productResponses,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	var product ep.Product
	productID := c.Param("id")

	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	var productImage ep.ProductImage
	productImages, err := h.productUseCase.GetProductImageURLById(productID, &productImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product images",
			"error":   err,
		})
	}

	var imageURLs []string
	for _, image := range productImages {
		imageURLs = append(imageURLs, image.Product_image_url)
	}

	productResponse := ep.ProductResponse{
		Product_id:        product.ID,
		Name:              product.Name,
		Category:          product.Product_Category.Category,
		Stock:             product.Stock,
		Price:             product.Price,
		Status:            product.Status,
		Rating:            product.Rating,
		Description:       product.Product_Description.Description,
		Product_image_url: imageURLs,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get product data by id",
		"products": productResponse,
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productID := c.Param("id")
	var req ep.ProductRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to bind data",
			"error":   err,
		})
	}

	status := "tersedia"
	if req.Stock == 0 {
		status = "habis"
	}

	req = ep.ProductRequest{
		Product_category_id: req.Product_category_id,
		Name:                req.Name,
		Stock:               req.Stock,
		Price:               req.Price,
		Description:         req.Description,
		Status:              status,
		Product_image_url:   req.Product_image_url,
	}

	var product ep.Product
	product, err = h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	err = h.productUseCase.UpdateProduct(productID, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update data",
			"error":   err,
		})
	}

	err = h.productUseCase.UpdateProductDescription(fmt.Sprint(product.Product_description_id), req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update product description",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update product",
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productID := c.Param("id")

	var product ep.Product
	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	var productDescription ep.ProductDescription
	err = h.productUseCase.DeleteProductDescription(fmt.Sprint(product.Product_description_id), &productDescription)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete product description",
			"error":   err,
		})
	}

	var productImages []ep.ProductImage
	err = h.productUseCase.DeleteProductImage(productID, &productImages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete product image",
			"error":   err,
		})
	}

	err = h.productUseCase.DeleteProduct(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete product",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully delete product",
	})
}

func (h *ProductHandler) SearchProduct(c echo.Context) error {
	param := c.QueryParam("param")

	switch param {
	case "id":
		var product ep.Product
		productID := c.QueryParam("id")
		product, err := h.productUseCase.SearchProductByID(productID, &product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product",
				"error":   err,
			})
		}

		var productImage ep.ProductImage
		productImages, err := h.productUseCase.GetProductImageURLById(productID, &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Product_Description.Description,
			Product_image_url: imageURLs,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product data by id",
			"products": productResponse,
		})
	case "name":
		var products []ep.Product
		name := c.QueryParam("name")
		products, err := h.productUseCase.SearchProductByName(name, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product",
				"error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Failed to get product images",
					"error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.Product_image_url)
			}

			productResponse := ep.ProductResponse{
				Product_id:        product.ID,
				Name:              product.Name,
				Category:          product.Product_Category.Category,
				Stock:             product.Stock,
				Price:             product.Price,
				Status:            product.Status,
				Rating:            product.Rating,
				Description:       product.Product_Description.Description,
				Product_image_url: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Product not found",
				"product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product by name",
			"products": productResponses,
		})
	case "category":
		var products []ep.Product
		category := c.QueryParam("category")
		products, err := h.productUseCase.SearchProductByCategory(category, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product by category",
				"error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Failed to get product images",
					"error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.Product_image_url)
			}

			productResponse := ep.ProductResponse{
				Product_id:        product.ID,
				Name:              product.Name,
				Category:          product.Product_Category.Category,
				Stock:             product.Stock,
				Price:             product.Price,
				Status:            product.Status,
				Rating:            product.Rating,
				Description:       product.Product_Description.Description,
				Product_image_url: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Product not found",
				"product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product data by category",
			"products": productResponses,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Invalid search parameter",
	})
}

func (h *ProductHandler) FilterProductByStatus(c echo.Context) error {
	status := c.QueryParam("status")

	var product []ep.Product
	products, err := h.productUseCase.FilterProductByStatus(status, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to filter product",
			"error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Product_Description.Description,
			Product_image_url: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get product based on stock status",
		"products": productResponses,
	})
}

func (h *ProductHandler) ImportProductFromCSV(c echo.Context) error {
	// file, err := c.FormFile("csvFile")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Failed to upload CSV file",
	// 		"error":   err,
	// 	})
	// }

	// src, err := file.Open()
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "Failed to open CSV file",
	// 		"error":   err,
	// 	})
	// }
	// defer src.Close()

	// reader := csv.NewReader(src)

	// records, err := reader.ReadAll()
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "Failed to read CSV file",
	// 		"error":   err,
	// 	})
	// }

	// totalWorker := 100
	// wg := &sync.WaitGroup{}

	// for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
	// 	go func(wg *sync.WaitGroup) {
	// 		defer wg.Done()

	// 		header := make(map[string]int)
	// 		for i, columnName := range records[0] {
	// 			header[columnName] = i
	// 		}

	// 		for _, record := range records[1:] {
	// 			product_category_id, _ := strconv.Atoi(record[header["product_category_id"]])
	// 			name := record[header["name"]]
	// 			// stock, _ := strconv.Atoi(record[header["stock"]])
	// 			price, _ := strconv.ParseFloat(record[header["price"]], 64)
	// 			description := record[header["description"]]
	// 			productImageURLs := strings.Split(record[header["product_image_url"]], ",")

	// 			productDescription := ep.ProductDescription{
	// 				Description: description,
	// 			}
	// 			err = h.productUseCase.CreateProductDescription(&productDescription)
	// 			if err != nil {
	// 				// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 				// 	"message": "Failed to create product description",
	// 				// 	"error":   err,
	// 				// })
	// 			}

	// 			status := "tersedia"
	// 			if stock == 0 {
	// 				status = "habis"
	// 			}

	// 			product := ep.Product{
	// 				Product_category_id:    uint(product_category_id),
	// 				Product_description_id: productDescription.ID,
	// 				Name:                   name,
	// 				Stock:                  stock,
	// 				Price:                  price,
	// 				Status:                 status,
	// 				Rating:                 0.00,
	// 			}
	// 			err = h.productUseCase.CreateProduct(&product)
	// 			if err != nil {

	// 			}

	// 			for _, url := range productImageURLs {
	// 				productImage := ep.ProductImage{
	// 					Product_id:        product.ID,
	// 					Product_image_url: url,
	// 				}
	// 				err = h.productUseCase.CreateProductImage(&productImage)
	// 				if err != nil {

	// 				}
	// 			}
	// 		}
	// 	}(wg)
	// }

	// wg.Add(totalWorker)
	// wg.Wait()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product imported successfully",
		// "products": productResponses,
	})
}
