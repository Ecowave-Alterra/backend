package product

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	productCategoryIDstr := c.FormValue("ProductCategoryId")
	productCategoryID, err := strconv.ParseUint(productCategoryIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid product category ID",
			"Error":   err,
		})
	}
	name := c.FormValue("Name")
	stockStr := c.FormValue("Stock")
	stock, err := strconv.ParseUint(stockStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid stock",
			"Error":   err,
		})
	}
	priceStr := c.FormValue("Price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid price",
			"Error":   err,
		})
	}
	description := c.FormValue("Description")

	status := "tersedia"
	if stock == 0 {
		status = "habis"
	}

	product := ep.Product{
		ProductCategoryId: uint(productCategoryID),
		Name:              name,
		Stock:             uint(stock),
		Price:             price,
		Status:            status,
		Rating:            0.00,
		Description:       description,
	}
	err = h.productUseCase.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create product",
			"Error":   err,
		})
	}

	cloudstorage.Folder = "img/products/"
	for i := 1; i <= 10; i++ {
		fileHeader, err := c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, Anda harus mengunggah foto",
				})
			}
			fileExtension := filepath.Ext(fileHeader.Filename)
			allowedExtensions := map[string]bool{
				".png":  true,
				".jpeg": true,
				".jpg":  true,
			}
			if !allowedExtensions[fileExtension] {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
				})
			}
			maxFileSize := 4 * 1024 * 1024
			fileSize := fileHeader.Size
			if fileSize > int64(maxFileSize) {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
				})
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			log.Println(PhotoUrl)

			productImage := ep.ProductImage{
				ProductId:       product.ID,
				ProductImageUrl: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Failed to create product image",
					"Error":   err,
				})
			}

		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menambahkan produk",
	})
}

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProduct(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product images",
				"Error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.ProductImageUrl)
		}

		productResponse := ep.ProductResponse{
			ProductID:       product.ProductID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           product.Stock,
			Price:           product.Price,
			Status:          product.Status,
			Rating:          product.Rating,
			Description:     product.Description,
			ProductImageUrl: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get all product datas",
		"Products": productResponses,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	var product ep.Product
	productID := c.Param("id")

	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product",
			"Error":   err,
		})
	}

	var productImage ep.ProductImage
	productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product images",
			"Error":   err,
		})
	}

	var imageURLs []string
	for _, image := range productImages {
		imageURLs = append(imageURLs, image.ProductImageUrl)
	}

	productResponse := ep.ProductResponse{
		ProductID:       product.ProductID,
		Name:            product.Name,
		Category:        product.ProductCategory.Category,
		Stock:           product.Stock,
		Price:           product.Price,
		Status:          product.Status,
		Rating:          product.Rating,
		Description:     product.Description,
		ProductImageUrl: imageURLs,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get product data by id",
		"Products": productResponse,
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("id")
	var req ep.ProductRequest

	productCategoryIDstr := c.FormValue("ProductCategoryId")
	productCategoryID, err := strconv.ParseUint(productCategoryIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid product category ID",
			"Error":   err,
		})
	}
	name := c.FormValue("Name")
	stockStr := c.FormValue("Stock")
	stock, err := strconv.ParseUint(stockStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid stock",
			"Error":   err,
		})
	}
	priceStr := c.FormValue("Price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid price",
			"Error":   err,
		})
	}
	description := c.FormValue("Description")

	status := "tersedia"
	if stock == 0 {
		status = "habis"
	}

	req = ep.ProductRequest{
		ProductCategoryId: uint(productCategoryID),
		Name:              name,
		Price:             price,
		Description:       description,
		Status:            status,
	}

	var product ep.Product
	product, err = h.productUseCase.GetProductByID(productId, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product",
			"Error":   err,
		})
	}

	err = h.productUseCase.UpdateProduct(productId, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to update data",
			"Error":   err,
		})
	}

	err = h.productUseCase.UpdateProductStock(productId, req.Stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to update product stock",
			"Error":   err,
		})
	}

	var productImages []ep.ProductImage
	err = h.productUseCase.DeleteProductImage(fmt.Sprint(product.ID), &productImages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to delete product images",
			"Error":   err,
		})
	}

	cloudstorage.Folder = "img/products/"
	for i := 1; i <= 5; i++ {
		fileHeader, err := c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			filename, _ := cloudstorage.GetFileName(fileHeader.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal mendapatkan nama file",
				})
			}
			err = cloudstorage.DeleteImage(filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal menghapus file pada cloud storage",
				})
			}
			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			productImage := ep.ProductImage{
				ProductId:       product.ID,
				ProductImageUrl: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Failed to create product image",
					"Error":   err,
				})
			}

		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah produk",
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	var product ep.Product
	product, err := h.productUseCase.GetProductByID(productId, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product",
			"Error":   err,
		})
	}

	var productImage ep.ProductImage
	productImages, _ := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)

	for _, image := range productImages {
		filename, err := cloudstorage.GetFileName(image.ProductImageUrl)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal mendapatkan nama file",
			})
		}
		err = cloudstorage.DeleteImage(filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal menghapus file pada cloud storage",
			})
		}

		err = h.productUseCase.DeleteProductImageByID(strconv.Itoa(int(image.ID)), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to delete product image",
			})
		}
	}

	err = h.productUseCase.DeleteProduct(productId, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to delete product",
			"Error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus produk",
	})
}

func (h *ProductHandler) SearchProduct(c echo.Context) error {
	param := c.QueryParam("param")

	switch param {
	case "id":
		var product ep.Product
		itemId := c.QueryParam("id")
		product, err := h.productUseCase.SearchProductByID(itemId, &product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product",
				"Error":   err,
			})
		}

		var productImage ep.ProductImage
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product images",
				"Error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.ProductImageUrl)
		}

		productResponse := ep.ProductResponse{
			ProductID:       product.ProductID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           product.Stock,
			Price:           product.Price,
			Status:          product.Status,
			Rating:          product.Rating,
			Description:     product.Description,
			ProductImageUrl: imageURLs,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message":  "Successfully get product data by id",
			"Products": productResponse,
		})
	case "name":
		var products []ep.Product
		name := c.QueryParam("name")
		products, err := h.productUseCase.SearchProductByName(name, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product",
				"Error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product images",
					"Error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.ProductImageUrl)
			}

			productResponse := ep.ProductResponse{
				ProductID:       product.ProductID,
				Name:            product.Name,
				Category:        product.ProductCategory.Category,
				Stock:           product.Stock,
				Price:           product.Price,
				Status:          product.Status,
				Rating:          product.Rating,
				Description:     product.Description,
				ProductImageUrl: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Product not found",
				"Product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message":  "Successfully get product by name",
			"Products": productResponses,
		})
	case "category":
		var products []ep.Product
		category := c.QueryParam("category")
		products, err := h.productUseCase.SearchProductByCategory(category, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product by category",
				"Error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product images",
					"Error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.ProductImageUrl)
			}

			productResponse := ep.ProductResponse{
				ProductID:       product.ProductID,
				Name:            product.Name,
				Category:        product.ProductCategory.Category,
				Stock:           product.Stock,
				Price:           product.Price,
				Status:          product.Status,
				Rating:          product.Rating,
				Description:     product.Description,
				ProductImageUrl: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Message": "Product not found",
				"Product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message":  "Successfully get product data by category",
			"Products": productResponses,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"Message": "Invalid search parameter",
	})
}

func (h *ProductHandler) FilterProductByStatus(c echo.Context) error {
	status := c.QueryParam("status")

	var product []ep.Product
	products, err := h.productUseCase.FilterProductByStatus(status, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to filter product",
			"Error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product images",
				"Error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.ProductImageUrl)
		}

		productResponse := ep.ProductResponse{
			ProductID:       product.ProductID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           product.Stock,
			Price:           product.Price,
			Status:          product.Status,
			Rating:          product.Rating,
			Description:     product.Description,
			ProductImageUrl: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get product based on stock status",
		"Products": productResponses,
	})
}

func (h *ProductHandler) DownloadCSVFile(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProduct(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	file, err := os.Create("product-data.csv")
	defer file.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to create csv file",
			"Error":   err,
		})
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	csvHeader := []string{"Product_id", "Name", "Category", "Stock", "Price", "Status", "Rating", "Description", "ProductImageUrl"}
	var csvData [][]string
	csvData = append(csvData, csvHeader)

	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product images",
				"Error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.ProductImageUrl)
		}

		record := []string{
			product.ProductID,
			product.Name,
			product.ProductCategory.Category,
			strconv.Itoa(int(product.Stock)),
			strconv.FormatFloat(product.Price, 'f', -1, 64),
			product.Status,
			strconv.FormatFloat(product.Rating, 'f', -1, 64),
			product.Description,
			strings.Join(imageURLs, ", "),
		}

		csvData = append(csvData, record)
	}

	w.WriteAll(csvData)

	if err := w.Error(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Failed to write CSV file",
			"Error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully generate CSV file",
	})
}
