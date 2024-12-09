package service

import (
	"fmt"
	"strconv"

	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"
	"golang.org/x/crypto/bcrypt"
)

type User = model.User

type UserService interface {
	Create(user User) error
	GetByID(id string) (User, error)
	Update(user User) error
	Delete(id string) error
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	Authenticate(username, password string) (User, error)
	ChangePassword(id string, oldPassword, newPassword string) error
	ResetPassword(id uint, newPassword string) error
	ChangeEmail(id string, newEmail string) error
	ForgotPassword(email string) (User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return *userService{repo: repo}
}

func (s *userService) Create(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(&user)
}

func (s *userService) GetByID(id string) (*User, error) {
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return &User{}, fmt.Errorf("invalid ID format")
	}

	return s.repo.GetByID(uint(userID)) // pastikan mengubah ke uint
}

func (s *userService) Update(user User) error {
	return s.repo.Update(&user)
}

func (s *userService) Delete(id string) error {
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}

	return s.repo.Delete(uint(userID)) // pastikan mengubah ke uint
}

func (s *userService) FindByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) FindByUsername(username string) (*User, error) {
	return s.repo.FindByUsername(username)
}

func (s *userService) Authenticate(username, password string) (*User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return &User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &User{}, fmt.Errorf("invalid username or password")
	}

	return user, nil
}

func (s *userService) ChangePassword(id string, oldPassword, newPassword string) error {
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}

	user, err := s.repo.GetByID(uint(userID))
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return fmt.Errorf("invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.Update(user)
}

func (s *userService) ResetPassword(id uint, newPassword string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.Update(user)
}

func (s *userService) ForgotPassword(email string) error {
	// Handle forgot password logic
	return s.repo.ForgotPassword(email)
}

func (s *userService) sendPasswordResetEmailEmail(id string, newEmail string) error {
	// Handle sending password reset email logic
	return nil
}
