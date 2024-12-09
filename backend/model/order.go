package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint      `gorm:"primaryKey","required","unique" json:"id"`
	UserID    uint      `json:userID`
	User      []User    `gorm:"foreignKey:"UserID"`
	OrderDate time.Time `json:"order_date"`
	Status    int       `json:"status"` //created, prosessed, complete, cancelled
	Product   []Product `gorm:"foreignKey:ProductID"`
	Count     int64     `json:"count"`
}

type OrderItem struct {
	ID        uint           `gorm:"primaryKey","required","unique" json:"id"`
	OrderID   uint           `json:order_id`
	Order     []Order        `gorm:"foreignKey:OrderID"`
	ProductID uint           `json:product_id`
	Product   []Product      `gorm:"foreignKey:ProductID"`
	Quantity  int            `json:"quantity"`
	Price     float64        `json:"price"`
	Discount  float64        `json:"discount"`
	Subtotal  float64        `json:"subtotal"`
	Total     float64        `json:"total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    int            `json:"status"` //pending, shipped, delivered, returned
}
