package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type Stock = model.HistoryStock

type stockRepository interface {
	GetAllStocks() ([]Stock, error)
	GetStockByID(id int) (*Stock, error)
	CreateStock(stock *Stock) error
	UpdateStock(stock *Stock) error
	DeleteStock(id int) error
	GetStockByProductID(productId int) ([]Stock, error)
}

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (repo *StockRepository) GetAllStocks() ([]Stock, error) {
	var stocks []Stock
	err := repo.db.Find(&stocks).Error
	return stocks, err
}

func (repo *StockRepository) GetStockByID(id int) (*Stock, error) {
	var stock Stock
	err := repo.db.Where("id = ?", id).First(&stock).Error
	return &stock, err
}

func (repo *StockRepository) CreateStock(stock *Stock) error {
	return repo.db.Create(stock).Error
}

func (repo *StockRepository) UpdateStock(stock *Stock) error {
	return repo.db.Save(stock).Error
}

func (repo *StockRepository) DeleteStock(id int) error {
	return repo.db.Delete(&Stock{StockID: id}).Error
}

func (repo *StockRepository) GetStockByProductID(productId int) ([]Stock, error) {
	var stocks []Stock
	err := repo.db.Where("product_id = ?", productId).Find(&stocks).Error
	return stocks, err
}
