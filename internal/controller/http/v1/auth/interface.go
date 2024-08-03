package auth_router

import (
	auth_model "contact_center/internal/domain/auth/model"
	"context"
)

type UseCase interface {
	Register(ctx context.Context, model auth_model.CreateUserRequest) (int64, error)
	GetToken(ctx context.Context, model auth_model.GetTokenRequest) (string, error)
}
