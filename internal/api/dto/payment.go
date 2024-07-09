package dto

import (
	"github.com/google/uuid"
	"payment-payments-api/internal/models"
	"payment-payments-api/internal/models/enums"
	"time"
)

func MapPaymenToPaymentResponse(model *models.Payment) PaymentResponse {
	return PaymentResponse{
		PaymentID:     model.ID,
		TransactionID: model.TransactionID,
		RefundID:      model.RefundID,
		UserID:        model.UserID,
		MerchantID:    model.MerchantID,
		Msg:           model.Msg,
		Amount:        model.Amount,
		Currency:      model.Currency,
		Status:        model.Status,
		Merchant:      model.Merchant,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

type PaymentResponse struct {
	PaymentID     uuid.UUID           `json:"PaymentId"`
	TransactionID string              `json:"transactionId"`
	RefundID      string              `json:"refundID"`
	UserID        string              `json:"userId"`
	MerchantID    string              `json:"merchantId"`
	Msg           string              `json:"msg"`
	Amount        float64             `json:"amount"`
	Currency      string              `json:"currency"`
	Status        enums.PaymentStatus `json:"status"`
	Merchant      string              `json:"merchant"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}

type PaymentRequest struct {
	CardID      string  `json:"cardId"`
	CVC         string  `json:"cvc"`
	ExpiredDate string  `json:"expiredDate"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Merchant    string  `json:"merchant"`
	UserID      string  `json:"userId"`
	MerchantID  string  `json:"merchantId"`
}
