package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type Category = model.Category
// type Product = model.Product

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id uint) (*Category, error) {
	var category Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return &category, err
}

// Fungsi untuk mencari kategori berdasarkan ID produk
func (r *CategoryRepository) FindByProductID(productID uint) (*Category, error) {
	var product Product
	err := r.db.Preload("Category").Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product.Category, nil
}
