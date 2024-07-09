package dto

import (
	"payment-payments-api/internal/api/dto"
	"payment-payments-api/internal/models/enums"
)

func MapPaymentRequestToPaymentRequest(dto dto.PaymentRequest) PaymentRequest {
	return PaymentRequest{
		CardID:      dto.CardID,
		CVC:         dto.CVC,
		ExpiredDate: dto.ExpiredDate,
		Amount:      dto.Amount,
		Currency:    dto.Currency,
		Merchant:    dto.Merchant,
	}
}

type PaymentRequest struct {
	PaymentID     string              `json:"paymentId"`
	TransactionID string              `json:"transactionId"`
	Status        enums.PaymentStatus `json:"status"`
	CardID        string              `json:"cardId"`
	CVC           string              `json:"cvc"`
	ExpiredDate   string              `json:"expiredDate"`
	Amount        float64             `json:"amount"`
	Type          enums.PaymentType   `json:"type"`
	Currency      string              `json:"currency"`
	Merchant      string              `json:"merchant"`
}

type PaymentResponse struct {
	PaymentID     string `json:"paymentID"`
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
	Msg           string `json:"msg"`
	RefundID      string `json:"refundID"`
}
