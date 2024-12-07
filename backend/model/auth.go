package model

type Admin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
