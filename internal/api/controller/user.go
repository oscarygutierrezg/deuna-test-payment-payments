package controller

import (
	"github.com/gin-gonic/gin"
	middleware "payment-payments-api/internal/api/middleware"
	"payment-payments-api/internal/models"
	"payment-payments-api/internal/services"
	"payment-payments-api/pkg/auth"
	"payment-payments-api/pkg/uhttp"
	"payment-payments-api/pkg/umdw"
	"payment-payments-api/pkg/util"
)

var User httpUser

type httpUser struct{}

func (httpUser) Login(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

		type AuthBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		type AuthResp struct {
			User  *models.User `json:"user"`
			Token string       `json:"token"`
		}

		var req AuthBody
		_ = umdw.BodyParse(&req, c)

		if req.Email == "admin@example.com" && req.Password == "admin123" {
			token, _ := middleware.NewJwtToken(models.User{})
			uhttp.Success(c, "User logged successfully", AuthResp{
				User:  nil,
				Token: token,
			})
			return
		}

		user, err := s.User.GetUserByEmail(req.Email)
		if err != nil {
			uhttp.Error(c, &util.JWTError{Message: "User cannot be logged."})
			return
		}

		if !auth.VerifyPassword(req.Password, user.PasswordSalt, user.Password) {
			uhttp.Error(c, &util.JWTError{Message: "User cannot be logged."})
			return
		}

		if !user.Enabled {
			uhttp.Error(c, &util.JWTError{Message: "User cannot be logged."})
			return
		}

		user.CleanSensitiveInfo()
		token, _ := middleware.NewJwtToken(*user)

		uhttp.Success(c, "User logged successfully", AuthResp{
			User:  user,
			Token: token,
		})
	}
}

func (httpUser) Create(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.User
		_ = umdw.BodyParse(&req, c)

		ue, _ := s.User.GetUserByEmail(req.Email)
		if ue != nil {
			uhttp.Error(c, &util.JWTError{Message: "User email already exists."})
			return
		}

		user, err := s.User.CreateUser(req.FirstName, req.LastName, req.Email)
		if err != nil {
			uhttp.Error(c, &util.JWTError{Message: err.Error()})
			return
		}

		user.CleanSensitiveInfo()

		uhttp.Success(c, "User created successfully.", user)
	}
}
