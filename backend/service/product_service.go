package service

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"
)

type Product = model.Product

type ProductService interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
	GetAllProducts() ([]Product, error)
	GetProductById(id int) (*Product, error)
	GetProductByUserId(userId uint) ([]Product, error)
	AssignCategoryToProduct(category *Category) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product *Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

func (s *productService) GetAllProducts() ([]Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) GetProductById(id int) (*Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) GetProductByUserId(userId uint) ([]Product, error) {
	return s.repo.GetProductByUserId(userId)
}

func (s *productService) AssignCategoryToProduct(category *Category) error {
	return s.repo.AssignCategoryToProduct(category)
}
