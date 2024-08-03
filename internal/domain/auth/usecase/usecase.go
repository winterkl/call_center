package auth_usecase

import (
	auth_model "contact_center/internal/domain/auth/model"
	"context"
)

type UseCase struct {
	api Api
}

func New(api Api) *UseCase {
	return &UseCase{
		api: api,
	}
}

func (uc *UseCase) Register(ctx context.Context, model auth_model.CreateUserRequest) (int64, error) {
	authModel, err := uc.api.Register(ctx, &model)
	if err != nil {
		return 0, err
	}
	return authModel.ID, nil
}
func (uc *UseCase) GetToken(ctx context.Context, model auth_model.GetTokenRequest) (string, error) {
	authModel, err := uc.api.GetToken(ctx, &model)
	if err != nil {
		return "", err
	}
	return authModel.Token, nil
}
