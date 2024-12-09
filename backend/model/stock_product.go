package model

import "errors"

type HistoryStock struct {
	StockID   int       `gorm:"primaryKey" json:"stock_id"`
	ProductID []Product `gorm:"not null" json:"foreignKey:product_id"`
	Quantity  int       `json:"quantity"`
}

// Validate memeriksa apakah data stock valid
func (s *HistoryStock) Validate() error {
	if s.ProductID == 0 {
		return errors.New("product ID is required")
	}
	if s.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if s.StockDate == "" {
		return errors.New("stock date is required")
	}
	// Tambahkan validasi lainnya jika perlu
	return nil
}
