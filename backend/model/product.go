package model

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey","required","unique" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	CreatedAt   string         `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   string         `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Images      byte           `json:"image"`
	Description string         `json:"description"`
	Variants    []Variants     `gorm:"foreignKey:ProductID"`
	Stock       []HistoryStock `gorm:"foreignKey:ProductID"`
}

type Variants struct {
	ID        uint      `gorm:"primaryKey","required","unique" json:"id"`
	ProductID []Product `gorm:"not null" json:"foreignKey:product_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Value     float64   `gorm:"type:decimal(10,2);not null" json:"value"`
}
