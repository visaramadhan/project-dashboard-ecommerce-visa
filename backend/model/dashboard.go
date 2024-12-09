package model

type Summary struct {
	TotalUsers int `json:"total_user"`
	TotalOrder int `json:"total_order"`
	TotalSales int `json:"total_sales"`
	TotalItems int `json:"total_items"`
}
