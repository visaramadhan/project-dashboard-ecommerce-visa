package repository

import "gorm.io/gorm"

type Repository struct {
	Auth     UserRepository
	Product  ProductRepository
	Category CategoryRepository
	Order    OrderRepository
	Stock    StockRepository
	Variants VariantRepository
	Summary  SummaryRepository
	Banner   BannerRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Auth:     *NewUserRepository(db),
		Product:  *NewProductRepository(db),
		Category: *NewCategoryRepository(db),
		Order:    *NewOrderRepository(db),
		Stock:    *NewStockRepository(db),
		Variants: *NewVariantRepository(db),
		Summary:  *NewSummaryRepository(db),
		Banner:   *NewBannerRepository(db),
	}
}
