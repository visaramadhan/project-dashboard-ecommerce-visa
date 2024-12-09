package service

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"
)

type Stock = model.HistoryStock

type StockService interface {
	GetAllStocks() ([]Stock, error)
	GetStockByID(id int) (*Stock, error)
	CreateStock(stock *Stock) error
	UpdateStock(stock *Stock) error
	DeleteStock(id int) error
	GetStockByProductID(productId int) ([]Stock, error)
}

type stockService struct {
	repo repository.StockRepository
}

// Constructor untuk membuat instance StockService
func NewStockService(repo repository.StockRepository) StockService {
	return &stockService{repo: repo}
}

// Mengambil semua data stock

func (s *stockService) GetAllStocks() ([]Stock, error) {
	return s.repo.GetAllStocks()
}

// Mengambil data stock berdasarkan ID

func (s *stockService) GetStockByID(id int) (*Stock, error) {
	return s.repo.GetStockByID(id)
}

// Menambahkan data stock

func (s *stockService) CreateStock(stock *Stock) error {
	if err := stock.Validate(); err != nil {
		return err
	}
	return s.repo.CreateStock(stock)
}

// Mengubah data stock

func (s *stockService) UpdateStock(stock *Stock) error {
	if err := stock.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateStock(stock)
}

// Menghapus data stock berdasarkan ID

func (s *stockService) DeleteStock(id int) error {
	return s.repo.DeleteStock(id)
}

// Mengambil data stock berdasarkan ID produk

func (s *stockService) GetStockByProductID(productId int) ([]Stock, error) {
	return s.repo.GetStockByProductID(productId)
}
