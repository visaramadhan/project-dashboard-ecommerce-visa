package repository

import (
	"errors"
	"fmt"
	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create saves a new user to the database
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update saves changes to an existing user in the database
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete removes a user from the database
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{ID: id}).Error
}

// FindByEmail retrieves a user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername retrieves a user by username
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Authenticate verifies the user's credentials
func (r *UserRepository) Authenticate(email, password string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Compare the password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

// ChangePassword changes the user's password
func (r *UserRepository) ChangePassword(id uint, oldPassword, newPassword string) error {
	// Find the user by ID
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("old password is incorrect")
	}

	// Hash new password and update
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return r.db.Save(&user).Error
}

// ResetPassword resets the user's password to a new one
func (r *UserRepository) ResetPassword(id uint, newPassword string) error {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return r.db.Save(&user).Error
}

// ForgotPassword handles forgot password logic (typically sending reset email)
func (r *UserRepository) ForgotPassword(email string) error {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return errors.New("email not found")
	}

	// Generate a password reset link (here we simulate it)
	resetLink := fmt.Sprintf("http://yourapp.com/reset-password?token=some_unique_token")

	// Send the reset password email
	return r.sendPasswordResetEmail(user, resetLink)
}

// VerifyEmail verifies the user's email using a token
func (r *UserRepository) VerifyEmail(token string) error {
	// In a real scenario, you'd verify the token. Simulating here.
	var user model.User
	err := r.db.Where("email_verification_token = ?", token).First(&user).Error
	if err != nil {
		return errors.New("invalid or expired token")
	}

	user.EmailVerified = true
	return r.db.Save(&user).Error
}

// ResendVerificationEmail resends the email verification link
func (r *UserRepository) ResendVerificationEmail(user *model.User) error {
	// Generate the verification link
	verificationLink := fmt.Sprintf("http://yourapp.com/verify-email?token=%s", user.EmailVerificationToken)

	// Send the verification email
	return r.sendVerificationEmail(user, verificationLink)
}

// SendPasswordResetEmail sends a password reset email
func (r *UserRepository) sendPasswordResetEmail(user model.User, resetLink string) error {
	// Setup the email configuration
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "your-email@example.com")
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Password Reset Request")
	mailer.SetBody("text/html", fmt.Sprintf("<p>Click <a href=\"%s\">here</a> to reset your password.</p>", resetLink))

	// SMTP server configuration
	dialer := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-email-password")

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}

// SendVerificationEmail sends a verification email to the user
func (r *UserRepository) sendVerificationEmail(user *model.User, verificationLink string) error {
	// Setup the email configuration
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "your-email@example.com")
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Email Verification Request")
	mailer.SetBody("text/html", fmt.Sprintf("<p>Click <a href=\"%s\">here</a> to verify your email address.</p>", verificationLink))

	// SMTP server configuration
	dialer := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-email-password")

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
