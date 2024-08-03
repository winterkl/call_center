package auth_api

import (
	auth_model "contact_center/internal/domain/auth/model"
	"context"
	authv1 "github.com/winterkl/auth_protobuf/gen/go/proto/auth"
)

type Api struct {
	authClient authv1.AuthClient
}

func New(authClient authv1.AuthClient) *Api {
	return &Api{
		authClient: authClient,
	}
}

func (a *Api) Register(ctx context.Context, model *auth_model.CreateUserRequest) (*auth_model.CreateUserResponse, error) {
	modelRequest := authv1.RegisterRequest{
		Login:    model.Login,
		Password: model.Password,
	}

	authModel, err := a.authClient.Register(ctx, &modelRequest)
	if err != nil {
		return nil, err
	}

	modelResponse := auth_model.CreateUserResponse{
		ID: authModel.Id,
	}

	return &modelResponse, nil
}
func (a *Api) GetToken(ctx context.Context, model *auth_model.GetTokenRequest) (*auth_model.GetTokenResponse, error) {
	modelRequest := authv1.GetTokenRequest{
		Login:    model.Login,
		Password: model.Password,
	}

	token, err := a.authClient.GetToken(ctx, &modelRequest)
	if err != nil {
		return nil, err
	}

	modelResponse := auth_model.GetTokenResponse{
		Token: token.Token,
	}
	return &modelResponse, nil
}
func (a *Api) ValidateToken(ctx context.Context, model *auth_model.ValidateTokenRequest) error {
	modelRequest := authv1.ValidateTokenRequest{
		Token: model.Token,
	}

	_, err := a.authClient.ValidateToken(ctx, &modelRequest)
	if err != nil {
		return err
	}
	return nil
}
