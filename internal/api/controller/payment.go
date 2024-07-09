package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"payment-payments-api/internal/api/dto"
	"payment-payments-api/internal/services"
	"payment-payments-api/pkg/uhttp"
	"payment-payments-api/pkg/umdw"
)

var Payment httpPayment

type httpPayment struct{}

func (httpPayment) Create(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.PaymentRequest
		_ = umdw.BodyParse(&req, c)

		res, err := s.Payment.CreatePayment(req)
		if err != nil {
			uhttp.Error(c, err)
			return
		}

		uhttp.Success(c, "Payment created successfully.", res)
	}
}

func (httpPayment) Refund(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RefundRequest
		_ = umdw.BodyParse(&req, c)

		res, err := s.Payment.RefundPayment(req)
		if err != nil {
			uhttp.Error(c, err)
			return
		}

		uhttp.Success(c, "Payment created successfully.", res)
	}
}

func (httpPayment) Get(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Params.ByName("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		payment, err := s.Payment.GetPaymentByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}

		uhttp.Success(c, "Payment created successfully.", payment)
	}
}
