package product

import (
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	cs "github.com/berrylradianh/ecowave-go/helper/customstatus"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	var products []ep.Product

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	products, total, err := h.productUseCase.GetAllProduct(&products, offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list produk",
			"Status":  http.StatusNotFound,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	var productResponses []ep.ProductResponse
	for _, product := range products {
		var imageURLs []string
		for _, image := range product.ProductImages {
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
		"Products":  productResponses,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	var product ep.Product
	productID := c.Param("id")

	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	var imageURLs []string
	for _, image := range product.ProductImages {
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
		"Products": productResponse,
		"Status":   http.StatusOK,
	})
}

func (h *ProductHandler) SearchProduct(c echo.Context) error {
	var products *[]ep.Product
	var err error

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	search := c.QueryParam("search")
	filter := c.QueryParam("filter")

	validParams := map[string]bool{"search": true, "filter": true, "page": true}
	for param := range c.QueryParams() {
		if !validParams[param] {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Masukkan parameter dengan benar",
				"Status":  http.StatusBadRequest,
			})
		}
	}

	products, total, err := h.productUseCase.SearchProduct(search, filter, offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	if len(*products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Produk yang anda cari tidak ditemukan",
			"Status":  http.StatusNotFound,
		})
	} else {
		var productResponses []ep.ProductResponse
		for _, product := range *products {
			var imageURLs []string
			for _, image := range product.ProductImages {
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
			"Products":  productResponses,
			"Page":      page,
			"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
			"Status":    http.StatusOK,
		})
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var product ep.Product
	productCategoryIDstr := c.FormValue("ProductCategoryId")
	if productCategoryIDstr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan product category ID",
			"Status":  http.StatusBadRequest,
		})
	} else {
		productCategoryID, err := strconv.ParseUint(productCategoryIDstr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid product category ID",
				"Status":  http.StatusBadRequest,
			})
		}
		product.ProductCategoryId = uint(productCategoryID)
	}

	name := c.FormValue("Name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan name",
			"Status":  http.StatusBadRequest,
		})
	} else {
		product.Name = name
	}

	stockStr := c.FormValue("Stock")
	if stockStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan stock",
			"Status":  http.StatusBadRequest,
		})
	} else {
		stock, err := strconv.ParseUint(stockStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
				"Status":  http.StatusBadRequest,
			})
		}
		product.Stock = uint(stock)
		if stock == 0 {
			product.Status = "habis"
		} else {
			product.Status = "tersedia"
		}
	}

	priceStr := c.FormValue("Price")
	if priceStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan price",
			"Status":  http.StatusBadRequest,
		})
	} else {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
				"Status":  http.StatusBadRequest,
			})
		}
		product.Price = float64(price)
	}

	description := c.FormValue("Description")
	if description == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan description",
			"Status":  http.StatusBadRequest,
		})
	} else {
		product.Description = description
	}

	err := h.productUseCase.CreateProduct(&product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	cloudstorage.Folder = "img/products/"
	photoUploaded := false
	for i := 1; i <= 5; i++ {
		fileHeader, err := c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			photoUploaded = true
			fileExtension := filepath.Ext(fileHeader.Filename)
			allowedExtensions := map[string]bool{
				".png":  true,
				".jpeg": true,
				".jpg":  true,
			}
			if !allowedExtensions[fileExtension] {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
					"Status":  http.StatusBadRequest,
				})
			}
			maxFileSize := 4 * 1024 * 1024
			fileSize := fileHeader.Size
			if fileSize > int64(maxFileSize) {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
					"Status":  http.StatusBadRequest,
				})
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			productImage := ep.ProductImage{
				ProductId:       product.ID,
				ProductImageUrl: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				code, msg := cs.CustomStatus(err.Error())
				return c.JSON(code, echo.Map{
					"Status":  code,
					"Message": msg,
				})
			}

		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	if !photoUploaded {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Mohon maaf anda harus mengunggah foto",
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Anda berhasil menambahkan produk",
		"Status":  http.StatusCreated,
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("id")
	var req ep.ProductRequest
	var product ep.Product

	productBefore, err := h.productUseCase.GetProductByID(productId, &product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	productCategoryIDstr := c.FormValue("ProductCategoryId")
	if productCategoryIDstr == "" {
		req.ProductCategoryId = productBefore.ProductCategoryId
	} else {
		productCategoryID, err := strconv.ParseUint(productCategoryIDstr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Invalid product category ID",
				"Status":  http.StatusBadRequest,
			})
		}
		req.ProductCategoryId = uint(productCategoryID)
	}

	name := c.FormValue("Name")
	if name == "" {
		req.Name = productBefore.Name
	} else {
		req.Name = name
	}

	stockStr := c.FormValue("Stock")
	if stockStr == "" {
		req.Stock = productBefore.Stock
	} else {
		stock, err := strconv.ParseUint(stockStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
				"Status":  http.StatusBadRequest,
			})
		}
		req.Stock = uint(stock)
		if stock == 0 {
			req.Status = "habis"
		} else {
			req.Status = "tersedia"
		}
	}

	priceStr := c.FormValue("Price")
	if priceStr == "" {
		req.Price = productBefore.Price
	} else {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": err,
				"Status":  http.StatusBadRequest,
			})
		}
		req.Price = float64(price)
	}

	description := c.FormValue("Description")
	if description == "" {
		req.Description = productBefore.Description
	} else {
		req.Description = description
	}

	err = h.productUseCase.UpdateProduct(productId, &req)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	err = h.productUseCase.UpdateProductStock(productId, req.Stock)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	var productImages []ep.ProductImage
	err = h.productUseCase.DeleteProductImage(fmt.Sprint(product.ID), &productImages)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	cloudstorage.Folder = "img/products/"
	var fileHeader *multipart.FileHeader
	for i := 1; i <= 5; i++ {
		fileHeader, _ = c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			fileExtension := filepath.Ext(fileHeader.Filename)
			allowedExtensions := map[string]bool{
				".png":  true,
				".jpeg": true,
				".jpg":  true,
			}
			if !allowedExtensions[fileExtension] {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
					"Status":  http.StatusBadRequest,
				})
			}
			maxFileSize := 4 * 1024 * 1024
			fileSize := fileHeader.Size
			if fileSize > int64(maxFileSize) {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
					"Status":  http.StatusBadRequest,
				})
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			productImage := ep.ProductImage{
				ProductId:       product.ID,
				ProductImageUrl: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				code, msg := cs.CustomStatus(err.Error())
				return c.JSON(code, echo.Map{
					"Status":  code,
					"Message": msg,
				})
			}
		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	if fileHeader != nil {
		for _, image := range productBefore.ProductImages {
			filename := cloudstorage.GetFileName(image.ProductImageUrl)
			err = cloudstorage.DeleteImage(filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Gagal menghapus file pada cloud storage",
					"Status":  http.StatusInternalServerError,
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah produk",
		"Status":  http.StatusOK,
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	var product ep.Product
	product, err := h.productUseCase.GetProductByID(productId, &product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	for _, image := range product.ProductImages {
		filename := cloudstorage.GetFileName(image.ProductImageUrl)
		err = cloudstorage.DeleteImage(filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Gagal menghapus file pada cloud storage",
				"Status":  http.StatusInternalServerError,
			})
		}

		err = h.productUseCase.DeleteProductImageByID(image.ID, &image)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}
	}

	err = h.productUseCase.DeleteProduct(productId, &product)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus produk",
		"Status":  http.StatusOK,
	})
}

func (h *ProductHandler) DownloadCSVFile(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProductNoPagination(&products)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	csvHeader := []string{"Product_id", "Name", "Category", "Stock", "Price", "Status", "Rating", "Description", "ProductImageUrl"}

	var productImage ep.ProductImage
	records := make([][]string, 0)
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
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

		records = append(records, record)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil membuat file CSV",
		"Status":  http.StatusOK,
		"Header":  csvHeader,
		"Records": records,
	})
}
