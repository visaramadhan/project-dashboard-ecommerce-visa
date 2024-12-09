package model


type HistoryStock struct {
	StockID    int     `gorm:"primaryKey" json:"stock_id"`
	ProductID []Product    `gorm:"not null" json:"foreignKey:product_id"`
}
