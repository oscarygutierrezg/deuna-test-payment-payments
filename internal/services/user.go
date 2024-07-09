package services

import (
	"github.com/google/uuid"
	"payment-payments-api/internal/models"
	"payment-payments-api/internal/repositories"
)

func NewUserService(r repositories.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

type userService struct {
	repo repositories.UserRepository
}

func (s *userService) CreateUser(firstName, lastName, email string) (*models.User, error) {
	user, err := models.NewUser(firstName, lastName, email)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserById(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetPaymentByEmail(email)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}
