package product

import (
	"time"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type Product struct {
	ProductId          string                 `gorm:"primarykey" json:"ProductId"`
	CreatedAt          time.Time              `json:"-"`
	UpdatedAt          *time.Time             `json:"-"`
	DeletedAt          *gorm.DeletedAt        `json:"-"`
	Name               string                 `validate:"required,max=10"`
	Stock              uint                   `validate:"required"`
	Price              float64                `validate:"required"`
	Status             string                 `validate:"required"`
	Weight             float64                `validate:"required"`
	Rating             float64                `validate:"required"`
	Description        string                 `validate:"required"`
	ProductCategoryId  uint                   `json:"-" validate:"required"`
	ProductCategory    ProductCategory        `gorm:"foreignKey:ProductCategoryId" json:"-"`
	ProductImages      []ProductImage         `gorm:"foreignKey:ProductId"`
	TransactionDetails []et.TransactionDetail `gorm:"foreignKey:ProductId"`
}

type ProductRequest struct {
	ProductCategoryId uint     `json:"productCategoryId" form:"productCategoryId"`
	Name              string   `json:"name" form:"name"`
	Stock             uint     `json:"stock" form:"stock"`
	Price             float64  `json:"price" form:"price"`
	Weight            float64  `validate:"required"`
	Description       string   `json:"description" form:"description"`
	Status            string   `json:"status" form:"status"`
	ProductImageUrl   []string `json:"productImageUrl" form:"productImageUrl"`
}

type ProductResponse struct {
	ProductId       string
	Name            string
	Category        string
	Stock           uint
	TotalOrders     uint    `json:"TotalOrders,omitempty"`
	TotalRevenue    float64 `json:"TotalRevenue,omitempty"`
	Weight          float64 `validate:"required"`
	Price           float64
	Status          string
	Rating          float64
	Description     string
	ProductImageUrl []string `json:"ProductImageUrl,omitempty"`
}
