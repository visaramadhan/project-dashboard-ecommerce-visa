package service

import (
	"fmt"
	"strconv"

	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/repository"
	"golang.org/x/crypto/bcrypt"
)

type User = model.User
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user model.User) error {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(&user)
}

func (s *UserService) FindByID(id string) (*model.User, error) {
	// Mengonversi ID string ke uint
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}

	// Menggunakan ID yang sudah dikonversi (uint)
	return s.repo.GetByID(uint(userID)) // pastikan mengubah ke uint
}
func (s *UserService) Update(user User) error {
	return s.repo.Update(&user)
}

func (s *UserService) Delete(id string) error {
	// Mengonversi ID string ke uint
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}

	// Panggil repositori dengan ID yang sudah dikonversi
	return s.repo.Delete(uint(userID)) // mengonversi ke uint
}

func (s *UserService) ListAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *UserService) FindByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Authenticate(email, password string) (*User, error) {
	return s.repo.Authenticate(email, password)
}

func (s *UserService) ChangePassword(id, oldPassword, newPassword string) error {
	// Mengonversi ID string ke uint
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}

	// Panggil repositori dengan ID yang sudah dikonversi
	return s.repo.ChangePassword(uint(userID), oldPassword, newPassword) // mengonversi ke uint
}

func (s *UserService) ResetPassword(id, newPassword string) error {
	// Reset password in repository
	userID, err := strconv.ParseUint(id, 10, 32) // 10: basis desimal, 32: ukuran bit
	if err != nil {
		return fmt.Errorf("invalid ID format")
	}
	// Panggil repositori dengan ID yang sudah dikonversi
	return s.repo.ResetPassword(uint(userID), newPassword)
}

func (s *UserService) ForgotPassword(email string) error {
	// Handle forgot password logic
	return s.repo.ForgotPassword(email)
}

func (s *UserService) VerifyEmail(token string) error {
	// Verify email using the token
	return s.repo.VerifyEmail(token)
}

func (s *UserService) ResendVerificationEmail(user *User) error {
	// Resend email verification link
	return s.repo.ResendVerificationEmail(user)
}
