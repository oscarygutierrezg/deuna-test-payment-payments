package api_route

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/internal/services"
)

func SetRoutes(r *gin.RouterGroup, s *services.Services) {
	paymentApi(r.Group("/payments"), s)
	userApi(r.Group("/users"), s)
}
