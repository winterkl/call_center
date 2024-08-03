package middleware

import (
	auth_model "contact_center/internal/domain/auth/model"
	"context"
)

type Validator interface {
	ValidateToken(ctx context.Context, model *auth_model.ValidateTokenRequest) error
}
