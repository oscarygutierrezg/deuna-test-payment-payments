package api_route

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/internal/api/controller"
	"payment-payments-api/internal/api/middleware"
	"payment-payments-api/internal/services"
)

func paymentApi(r *gin.RouterGroup, s *services.Services) {

	r.POST("",
		middleware.JwtValidation,
		middleware.Payment.CreateValidation,
		controller.Payment.Create(s),
	)

	r.GET("/:id",
		middleware.JwtValidation,
		controller.Payment.Get(s),
	)

	r.POST("/refund",
		middleware.JwtValidation,
		middleware.Refund.CreateValidation,
		controller.Payment.Refund(s),
	)
}
