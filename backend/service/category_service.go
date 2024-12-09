package service

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"
)

type Category = model.Category

type CategoryService interface {
	FindAll() ([]Category, error)
	FindByID(id uint) (*Category, error)
	FindByProductID(productID uint) (*Category, error)
	Update(category *Category) error
	Delete(id uint) error
	Create(category *Category) error}

type categoryService struct {
	repo repository.CategoryRepository
}

// Constructor untuk membuat instance CategoryService
func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

// FindAll menampilkan semua data kategori
 func (cs *categoryService) FindAll() ([]Category, error) {
    return cs.repo.FindAll()
}

// FindByID mencari data kategori berdasarkan ID
 func (cs *categoryService) FindByID(id uint) (*Category, error) {
    return cs.repo.FindByID(id)
}
 // FindByProductID mencari data kategori berdasarkan ID produk

 func (cs *categoryService) FindByProductID(productID uint) (*Category, error) {
    return cs.repo.FindByProductID(productID)
}

// Update mengupdate data kategori
 func (cs *categoryService) Update(category *Category) error {
    if err := cs.repo.Update(category); err != nil {
        return err
    }
    return nil
}

// Delete menghapus data kategori berdasarkan ID
 func (cs *categoryService) Delete(id uint) error {
    if err := cs.repo.Delete(id); err != nil {
        return err
    }
    return nil
}

// Create membuat data kategori baru
 func (cs *categoryService) Create(category *Category) error {
    if err := cs.repo.Create(category); err != nil {
        return err
    }
    return nil
}
