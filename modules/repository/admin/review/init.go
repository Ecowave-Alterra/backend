package review

import (
	"gorm.io/gorm"
)

type ReviewRepo interface {
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
