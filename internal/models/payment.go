package models

import (
	"github.com/google/uuid"
	"payment-payments-api/internal/models/enums"
	"time"
)

type Payment struct {
	ID            uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CardID        string              `json:"cardId"`
	TransactionID string              `json:"transactionId"`
	RefundID      string              `json:"refundID"`
	UserID        string              `json:"userId"`
	MerchantID    string              `json:"merchantId"`
	Amount        float64             `json:"amount"`
	Status        enums.PaymentStatus `json:"status"`
	Msg           string              `json:"msg"`
	Currency      string              `json:"currency"`
	Merchant      string              `json:"merchant"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}
