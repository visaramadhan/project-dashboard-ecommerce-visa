package repository

import "gorm.io/gorm"

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
    return &BannerRepository{db: db}
}