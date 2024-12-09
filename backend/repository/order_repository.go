package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type Order = model.Order
type OrderItem = model.OrderItem

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrdersByUserID(userID uint) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id =?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrderItemsByOrderID(orderID uint) ([]OrderItem, error) {
	var orderItems []OrderItem
	err := r.db.Where("order_id =?", orderID).Find(&orderItems).Error
	return orderItems, err
}

func (r *OrderRepository) GetOrderByID(orderID uint) (*Order, error) {
	var order Order
	err := r.db.Where("id =?", orderID).First(&order).Error
	return &order, err
}

func (r *OrderRepository) CreateOrder(order *Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) UpdateOrder(order *Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) DeleteOrder(order *Order) error {
	return r.db.Delete(order).Error
}

func (r *OrderRepository) DeleteOrderItemsByOrderID(orderID uint) error {
	return r.db.Where("order_id = ?", orderID).Delete(&OrderItem{}).Error
}

func (r *OrderRepository) DeleteOrdersByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&Order{}).Error
}

func (r *OrderRepository) TotalOrderAmountByUserID(userID uint) (float64, error) {
	var total float64
	err := r.db.Table("orders").Select("SUM(total_amount)").Where("user_id = ?", userID).Scan(&total).Error
	return total, err
}

func (r *OrderRepository) TotalOrderCountByUserID(userID uint) (int64, error) {
	var count int64 // Menggunakan int64 untuk menerima hasil dari Count
	err := r.db.Table("orders").Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *OrderRepository) TotalOrderAmountByDateRange(startDate, endDate string) (float64, error) {
	var total float64
	err := r.db.Table("orders").Select("SUM(total_amount)").Where("created_at BETWEEN ? AND ?", startDate, endDate).Scan(&total).Error
	return total, err
}

func (r *OrderRepository) TotalOrderCountByDateRange(startDate, endDate string) (int64, error) {
	var count int64 // Menggunakan int64 untuk menerima hasil dari Count
	err := r.db.Table("orders").Where("created_at BETWEEN ? AND ?", startDate, endDate).Count(&count).Error
	return count, err
}
