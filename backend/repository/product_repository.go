package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type Product = model.Product // Alias untuk model.Product

type productRepository interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
	GetAllProducts() ([]Product, error)
	GetProductById(id int) (*Product, error)
	GetProductByUserId(userId uint) ([]Product, error)
	AssignCategoryToProduct(category *Category) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) CreateProduct(product *Product) error {
	return pr.db.Create(product).Error
}

func (pr *ProductRepository) UpdateProduct(product *Product) error {
	return pr.db.Save(product).Error
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	return pr.db.Delete(&Product{}, id).Error
}

func (pr *ProductRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	return products, pr.db.Find(&products).Error
}

func (pr *ProductRepository) GetProductById(id int) (*Product, error) {
	var product Product
	return &product, pr.db.Where("id = ?", id).First(&product).Error
}

func (pr *ProductRepository) GetProductByUserId(userId uint) ([]Product, error) {
	var products []Product
	return products, pr.db.Where("user_id = ?", userId).Find(&products).Error
}

func (pr *ProductRepository) AssignCategoryToProduct(category *Category) error {
	// Append the category to the product's category association
	association := pr.db.Model(&Product{}).Association("Categories")
	if err := association.Append(category); err != nil {
		return err // Return the error if any occurs
	}
	return nil // Return nil if no error occurs
}
