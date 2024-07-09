package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"payment-payments-api/internal/models"
)

type PaymentRepository interface {
	CreatePayment(payment models.Payment) (models.Payment, error)
	GetPaymentByID(id uuid.UUID) (models.Payment, error)
	GetPaymentByTransactionID(transactionID string) (models.Payment, error)
	UpdatePayment(payment models.Payment) (models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment models.Payment) (models.Payment, error) {
	if err := r.db.Create(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *paymentRepository) GetPaymentByID(id uuid.UUID) (models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *paymentRepository) GetPaymentByTransactionID(transactionID string) (models.Payment, error) {
	var payment models.Payment
	if err := r.db.Where("transaction_id = ?", transactionID).First(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *paymentRepository) UpdatePayment(payment models.Payment) (models.Payment, error) {
	if err := r.db.Save(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}
