package services

import (
	"errors"
	"github.com/google/uuid"
	dtoApi "payment-payments-api/internal/api/dto"
	dtoKafka "payment-payments-api/internal/kafka/dto"
	"payment-payments-api/internal/kafka/producer"
	"payment-payments-api/internal/models"
	"payment-payments-api/internal/models/enums"
	"payment-payments-api/internal/repositories"
	"time"
)

var (
	PaymentAlreadyRefunded = errors.New("payment already refunded")
	PaymentNotFound        = errors.New("payment not found")
)

type PaymentService interface {
	CreatePayment(dto dtoApi.PaymentRequest) (models.Payment, error)
	RefundPayment(dto dtoApi.RefundRequest) (models.Payment, error)
	GetPaymentByID(id uuid.UUID) (dtoApi.PaymentResponse, error)
	UpdatePayment(payment dtoKafka.PaymentResponse) error
}

type paymentService struct {
	paymentProducer   producer.PaymentProducer
	paymentRepository repositories.PaymentRepository
}

func NewPaymentService(paymentRepository repositories.PaymentRepository,
	paymentProducer producer.PaymentProducer) *paymentService {
	return &paymentService{paymentRepository: paymentRepository,
		paymentProducer: paymentProducer,
	}
}

func (s *paymentService) CreatePayment(paymentRequest dtoApi.PaymentRequest) (models.Payment, error) {
	payment := models.Payment{
		CardID:     paymentRequest.CardID,
		Amount:     paymentRequest.Amount,
		Currency:   paymentRequest.Currency,
		Merchant:   paymentRequest.Merchant,
		MerchantID: paymentRequest.MerchantID,
		UserID:     paymentRequest.UserID,
		Status:     enums.Pending,
		CreatedAt:  time.Now(),
	}
	model, err := s.paymentRepository.CreatePayment(payment)
	if err != nil {
		return model, err
	}
	dto := dtoKafka.MapPaymentRequestToPaymentRequest(paymentRequest)
	dto.PaymentID = model.ID.String()
	dto.Status = model.Status
	dto.Type = enums.Payment
	err = s.paymentProducer.Produce(dto)
	return model, err
}

func (s *paymentService) GetPaymentByID(id uuid.UUID) (dtoApi.PaymentResponse, error) {
	dto := dtoApi.PaymentResponse{}
	model, err := s.paymentRepository.GetPaymentByID(id)
	if err != nil {
		return dto, PaymentNotFound
	}
	dto = dtoApi.MapPaymenToPaymentResponse(&model)
	return dto, nil
}

func (s *paymentService) RefundPayment(request dtoApi.RefundRequest) (models.Payment, error) {
	model, err := s.paymentRepository.GetPaymentByTransactionID(request.TransactionID)
	if err != nil {
		return model, PaymentNotFound
	}
	if model.Status == enums.Cancelled {
		return model, PaymentAlreadyRefunded
	}
	dto := dtoKafka.PaymentRequest{
		PaymentID:     model.ID.String(),
		TransactionID: request.TransactionID,
		Type:          enums.Refund,
		Status:        model.Status,
		Amount:        request.Amount,
		Currency:      request.Currency,
	}
	err = s.paymentProducer.Produce(dto)
	return model, err
}

func (s *paymentService) UpdatePayment(dto dtoKafka.PaymentResponse) error {
	id, err := uuid.Parse(dto.PaymentID)
	if err != nil {
		return err
	}
	payment, err := s.paymentRepository.GetPaymentByID(id)
	if err != nil {
		return err
	}
	payment.Status = enums.Parse(dto.Status)
	payment.TransactionID = dto.TransactionID
	payment.Msg = dto.Msg
	payment.RefundID = dto.RefundID
	payment.UpdatedAt = time.Now()
	_, err = s.paymentRepository.UpdatePayment(payment)
	return err
}
