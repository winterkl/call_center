package call_router

import (
	call_model "contact_center/internal/domain/call/model"
	"contact_center/pkg/postgres/utils/paginate"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, model call_model.CreateCallRequest) error
	Get(ctx context.Context, id int) (*call_model.GetCallResponse, error)
	GetList(ctx context.Context, filter call_model.CallFilter, pagination *paginate.Paginate) ([]call_model.GetCallResponse, error)
	Update(ctx context.Context, model call_model.UpdateCallRequest) error
	Delete(ctx context.Context, id int) error
}
