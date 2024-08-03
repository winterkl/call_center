package auth_usecase

import (
	auth_model "contact_center/internal/domain/auth/model"
	"context"
)

type Api interface {
	Register(ctx context.Context, model *auth_model.CreateUserRequest) (*auth_model.CreateUserResponse, error)
	GetToken(ctx context.Context, model *auth_model.GetTokenRequest) (*auth_model.GetTokenResponse, error)
}
