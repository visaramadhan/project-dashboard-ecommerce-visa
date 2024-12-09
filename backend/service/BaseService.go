package service

import "github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"

type Service struct {
	User     UserService
	Product  ProductService
	Category CategoryService
	Stock    StockService
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:     NewUserService(&repo.Auth),
		Product:  NewProductService(repo.Product),
		Category: NewCategoryService(repo.Category),
		Stock:    NewStockService(repo.Stock),
	}
}
