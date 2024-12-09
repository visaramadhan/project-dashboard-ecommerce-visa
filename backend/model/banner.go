package model

type Banner struct {
	ID      int       `gorm:"primaryKey" json:"id"`
	Title   string    `json:"title"`
	Image   string    `json:"image"`
	Product []Product `gorm:"foreignKey:ProductID"`
}
