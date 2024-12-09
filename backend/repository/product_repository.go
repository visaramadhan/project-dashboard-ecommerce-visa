package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type Product = model.Product // Alias untuk model.Product

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetProductById mendapatkan product berdasarkan id

func (r *ProductRepository) GetProductById(id uint) (*Product, error) {
	var product Product
	result := r.db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// GetProductByUserId mendapatkan product berdasarkan id user

func (r *ProductRepository) GetProductByUserId(userId uint) ([]Product, error) {
	var products []Product
	result := r.db.Where("user_id = ?", userId).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// CreateProduct membuat product baru

func (r *ProductRepository) CreateProduct(product *Product) error {
	result := r.db.Create(product)
	return result.Error
}

// UpdateProduct mengubah product

func (r *ProductRepository) UpdateProduct(product *Product) error {
	result := r.db.Save(product)
	return result.Error
}

// DeleteProduct menghapus product

func (r *ProductRepository) DeleteProduct(product *Product) error {
	result := r.db.Delete(product)
	return result.Error
}

// AssignCategoryToProduct menetapkan kategori ke produk
func (r *ProductRepository) AssignCategoryToProduct(productID, categoryID uint) error {
	// Cari produk berdasarkan ID
	var product Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return err // Produk tidak ditemukan
	}

	// Tetapkan kategori
	product.CategoryID = categoryID

	// Simpan perubahan
	if err := r.db.Save(&product).Error; err != nil {
		return err // Gagal menyimpan perubahan
	}

	return nil
}

// GetStockByProductID mendapatkan total stok berdasarkan ProductID
func (r *ProductRepository) GetStockByProductID(productID uint) (int, error) {
	var totalStock int
	err := r.db.Model(&model.HistoryStock{}).
		Where("product_id = ?", productID).
		Select("COALESCE(SUM(quantity), 0)"). // Menghitung total stok
		Scan(&totalStock).Error

	return totalStock, err
}
