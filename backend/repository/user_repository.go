package repository

import (
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"gorm.io/gorm"
)

type User = model.User // Alias untuk model.User

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetAll() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *UserRepository) Create(user *User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) Update(user *User) error {
	return u.db.Save(user).Error
}

func (u *UserRepository) Delete(id uint) error {
	return u.db.Delete(&User{ID: id}).Error
}

func (u *UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
