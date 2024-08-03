package call_router

import (
	"contact_center/internal/app_errors"
	"contact_center/internal/controller/http/response"
	"contact_center/pkg/postgres/utils/paginate"
	"errors"
	"github.com/gin-gonic/gin"
)

func handleCallError(ctx *gin.Context, err error) {
	var errInvalidPage *paginate.InvalidPage
	if errors.As(err, &errInvalidPage) {
		response.SendBadRequest(ctx, errInvalidPage.Error())
		return
	}
	var errCallNotFound *app_errors.CallNotFound
	if errors.As(err, &errCallNotFound) {
		response.SendNotFound(ctx, errCallNotFound.Error())
		return
	}
	var errInvalidDateRange *app_errors.BeginAfterEnd
	if errors.As(err, &errInvalidDateRange) {
		response.SendBadRequest(ctx, errInvalidDateRange.Error())
		return
	}
	response.SendInternalServerError(ctx, err)
}
