package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	models "payment-payments-api/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) error
	FindByID(id uuid.UUID) (*models.User, error)
	GetPaymentByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetPaymentByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}
