package auth_model

type CreateUserRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	ID int64
}

func NewGetUserResponse(id int64) CreateUserResponse {
	return CreateUserResponse{
		ID: id,
	}
}

type GetTokenRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetTokenResponse struct {
	Token string
}

func NewGetTokenResponse(token string) GetTokenResponse {
	return GetTokenResponse{
		Token: token,
	}
}

type ValidateTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type ValidateTokenResponse struct {
	IsValid bool
}

func NewValidateTokenResponse(isValid bool) ValidateTokenResponse {
	return ValidateTokenResponse{
		IsValid: isValid,
	}
}
