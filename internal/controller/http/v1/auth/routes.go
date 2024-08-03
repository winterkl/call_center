package auth_router

import (
	"contact_center/internal/controller/http/response"
	auth_model "contact_center/internal/domain/auth/model"
	"github.com/gin-gonic/gin"
)

type routes struct {
	uc UseCase
}

func New(handler *gin.RouterGroup, uc UseCase) {
	r := routes{uc: uc}
	h := handler.Group("/auth")
	{
		h.POST("", r.register)
		h.GET("", r.getToken)
	}
}

func (r *routes) register(ctx *gin.Context) {
	model := auth_model.CreateUserRequest{}
	if err := ctx.ShouldBind(&model); err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	userID, err := r.uc.Register(ctx, model)
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, userID)
}

func (r *routes) getToken(ctx *gin.Context) {
	model := auth_model.GetTokenRequest{}
	if err := ctx.ShouldBindJSON(&model); err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	token, err := r.uc.GetToken(ctx, model)
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, token)
}
