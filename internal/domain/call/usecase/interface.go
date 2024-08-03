package call_usecase

import (
	call_entity "contact_center/internal/domain/call/entity"
	call_model "contact_center/internal/domain/call/model"
	"contact_center/pkg/postgres/utils/paginate"
	"context"
)

type Repository interface {
	Create(ctx context.Context, call *call_entity.Call) error
	Get(ctx context.Context, id int) (*call_entity.Call, error)
	GetList(ctx context.Context, filter call_model.CallFilter, pagination *paginate.Paginate) ([]call_entity.Call, error)
	Update(ctx context.Context, call *call_entity.Call) error
	Delete(ctx context.Context, id int) error
}
