package middleware

import (
	"contact_center/internal/controller/http/response"
	auth_model "contact_center/internal/domain/auth/model"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidateToken(vd Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromHeader(ctx)
		if err != nil {
			response.SendUnauthorized(ctx, err.Error())
			return
		}
		model := auth_model.ValidateTokenRequest{
			Token: token,
		}

		//Отправляем запрос на валидацию в сервис авторизации
		if err = vd.ValidateToken(ctx, &model); err != nil {
			response.SendValidErrorRequest(ctx, err)
			return
		}

		ctx.Next()
	}
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not set")
	}

	bearerTokenParts := strings.Split(authHeader, "Bearer")
	if len(bearerTokenParts) < 2 {
		return "", errors.New("authorization header has wrong format")
	}

	token := strings.TrimSpace(bearerTokenParts[1])
	if token == "" {
		return "", errors.New("authorization token not set")
	}

	return token, nil
}
