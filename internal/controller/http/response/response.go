package response

import (
	"bytes"
	"contact_center/pkg/postgres/utils/paginate"
	"contact_center/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MIME_XLS = "application/vnd.ms-excel"
)

func failureResponse(ctx *gin.Context, status int, response *response.Response) {
	ctx.AbortWithStatusJSON(status, response)
}

func successResponse(ctx *gin.Context, status int, response *response.Response) {
	ctx.JSON(status, response)
}

func SendInternalServerError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	failureResponse(ctx, http.StatusInternalServerError,
		response.New("Internal Server Error"))
}

func SendBadRequest(ctx *gin.Context, text string) {
	failureResponse(ctx, http.StatusBadRequest, response.New(text))
}

func SendNotFound(ctx *gin.Context, text string) {
	failureResponse(ctx, http.StatusNotFound, response.New(text))
}

func SendUnauthorized(ctx *gin.Context, text string) {
	failureResponse(ctx, http.StatusUnauthorized, response.New(text))
}

func SendForbidden(ctx *gin.Context, text string) {
	failureResponse(ctx, http.StatusForbidden, response.New(text))
}

func SendOkRequest(ctx *gin.Context) {
	successResponse(ctx, http.StatusOK, response.New("success"))
}

func SendOkRequestWithData(ctx *gin.Context, data interface{}) {
	successResponse(ctx, http.StatusOK, response.New("success").SetData(data))
}

func SendOkRequestWithPaginationData(ctx *gin.Context, data interface{}, paginate *paginate.Paginate) {
	successResponse(ctx, http.StatusOK, response.New("success").SetData(data).SetPaginate(paginate))
}

func SendOkRequestWithFile(ctx *gin.Context, buffer *bytes.Buffer, fileName, contentType string) {
	// Устанавливаем заголовки для скачивания файла
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	// Отправляем файл
	ctx.Data(http.StatusOK, contentType, buffer.Bytes())
}

func SendValidErrorRequest(ctx *gin.Context, err error) {
	failureResponse(ctx, http.StatusBadRequest, response.New("Ошибка валидации").SetData(err.Error()))
}
