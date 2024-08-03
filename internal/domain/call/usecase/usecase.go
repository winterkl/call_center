package call_usecase

import (
	call_model "contact_center/internal/domain/call/model"
	"contact_center/pkg/postgres/utils/paginate"
	"context"
	"fmt"
)

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Create(ctx context.Context, model call_model.CreateCallRequest) error {
	if err := model.Validate(); err != nil {
		return err
	}

	call := model.GetEntity()
	if err := uc.repo.Create(ctx, call); err != nil {
		return fmt.Errorf("call_usecase -> Create -> repo.Create -> %w", err)
	}
	return nil
}
func (uc *UseCase) Get(ctx context.Context, id int) (*call_model.GetCallResponse, error) {
	call, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("call_usecase -> Get -> repo.Get -> %w", err)
	}
	return call_model.NewGetCallResponse(call), nil
}
func (uc *UseCase) GetList(ctx context.Context, filter call_model.CallFilter, pagination *paginate.Paginate) ([]call_model.GetCallResponse, error) {
	callList, err := uc.repo.GetList(ctx, filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("call_usecase -> GetList -> repo.GetList -> %w", err)
	}
	return call_model.GetCallResponseList(callList), nil
}
func (uc *UseCase) Update(ctx context.Context, model call_model.UpdateCallRequest) error {
	if err := model.Validate(); err != nil {
		return err
	}

	call := model.GetEntity()
	if err := uc.repo.Update(ctx, call); err != nil {
		return fmt.Errorf("call_usecase -> Update -> repo.Update -> %w", err)
	}
	return nil
}
func (uc *UseCase) Delete(ctx context.Context, id int) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("call_usecase -> Delete -> repo.Delete -> %w", err)
	}
	return nil
}
