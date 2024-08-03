package v1

import (
	"contact_center/internal/controller/http/middleware"
	auth_router "contact_center/internal/controller/http/v1/auth"
	call_router "contact_center/internal/controller/http/v1/call"
	"github.com/gin-gonic/gin"
)

type UC struct {
	CallUC call_router.UseCase
	AuthUC auth_router.UseCase
}

func New(handler *gin.Engine, uc UC, validator middleware.Validator) {
	h := handler.Group("v1")

	auth_router.New(h, uc.AuthUC)
	h.Use(middleware.ValidateToken(validator))
	{
		call_router.New(h, uc.CallUC)
	}
}
