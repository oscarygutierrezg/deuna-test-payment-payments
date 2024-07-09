package api_route

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/internal/api/controller"
	apimiddleware "payment-payments-api/internal/api/middleware"
	"payment-payments-api/internal/services"
)

func userApi(r *gin.RouterGroup, s *services.Services) {

	r.POST("/login",
		apimiddleware.User.LoginValidation,
		controller.User.Login(s),
	)

}
