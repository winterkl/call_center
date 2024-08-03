package call_router

import (
	"contact_center/internal/controller/http/response"
	call_model "contact_center/internal/domain/call/model"
	"contact_center/pkg/postgres/utils/paginate"
	"github.com/gin-gonic/gin"
	"strconv"
)

type routes struct {
	uc UseCase
}

func New(handler *gin.RouterGroup, uc UseCase) {
	r := routes{uc: uc}
	h := handler.Group("/call")
	{
		h.POST("", r.create)
		h.GET("/:id", r.get)
		h.GET("", r.getList)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}

func (r *routes) create(ctx *gin.Context) {
	model := call_model.CreateCallRequest{}
	if err := ctx.ShouldBind(&model); err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}

	if err := r.uc.Create(ctx, model); err != nil {
		handleCallError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func (r *routes) get(ctx *gin.Context) {
	callID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	call, err := r.uc.Get(ctx, callID)
	if err != nil {
		handleCallError(ctx, err)
		return
	}
	response.SendOkRequestWithData(ctx, call)
}

func (r *routes) getList(ctx *gin.Context) {
	pagination, err := paginate.New(ctx)
	if err != nil {
		handleCallError(ctx, err)
		return
	}
	filter := call_model.CallFilter{}
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = filter.Validate(); err != nil {
		handleCallError(ctx, err)
		return
	}

	callList, err := r.uc.GetList(ctx, filter, pagination)
	if err != nil {
		handleCallError(ctx, err)
		return
	}
	response.SendOkRequestWithPaginationData(ctx, callList, pagination)
}

func (r *routes) update(ctx *gin.Context) {
	callID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	model := call_model.UpdateCallRequest{ID: callID}
	if err = ctx.ShouldBindJSON(&model); err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.uc.Update(ctx, model); err != nil {
		handleCallError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}

func (r *routes) delete(ctx *gin.Context) {
	callID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendValidErrorRequest(ctx, err)
		return
	}
	if err = r.uc.Delete(ctx, callID); err != nil {
		handleCallError(ctx, err)
		return
	}
	response.SendOkRequest(ctx)
}
