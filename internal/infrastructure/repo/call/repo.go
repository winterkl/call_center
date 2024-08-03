package call_repo

import (
	"contact_center/internal/app_errors"
	call_entity "contact_center/internal/domain/call/entity"
	call_model "contact_center/internal/domain/call/model"
	"contact_center/pkg/postgres"
	"contact_center/pkg/postgres/utils/paginate"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
)

type Repo struct {
	db *postgres.Postgres
}

func New(db *postgres.Postgres) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, call *call_entity.Call) error {
	if err := r.db.NewInsert().
		Model(call).
		Scan(ctx); err != nil {
		return fmt.Errorf("NewInsert: %w", err)
	}
	return nil
}
func (r *Repo) Get(ctx context.Context, id int) (*call_entity.Call, error) {
	call := &call_entity.Call{}
	if err := r.db.NewSelect().
		Model(call).
		Order("caller_id").
		Relation("Agent").
		Relation("Status").
		Where("call.id = ?", id).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &app_errors.CallNotFound{}
		}
		return nil, fmt.Errorf("NewSelect: %w", err)
	}
	return call, nil
}
func (r *Repo) GetList(ctx context.Context, filter call_model.CallFilter, pagination *paginate.Paginate) ([]call_entity.Call, error) {
	var callList []call_entity.Call
	query := r.db.NewSelect().
		Model(&callList).
		Order("caller_id").
		Relation("Status").
		Relation("Agent").
		Where("call_start BETWEEN ? AND ?", filter.Begin, filter.End)

	if len(filter.CallerID) != 0 {
		query.Where("caller_id IN (?)", bun.In(filter.CallerID))
	}
	if len(filter.AgentID) != 0 {
		query.Where("agent_id IN (?)", bun.In(filter.AgentID))
	}
	if len(filter.CallStatus) != 0 {
		query.Where("call_status IN (?)", bun.In(filter.CallStatus))
	}

	if pagination != nil {
		if err := pagination.CalculatePagesCount(ctx, query); err != nil {
			return nil, err
		}
		if pagination.PageCounter < pagination.CurrentPage {
			return nil, nil
		}
		query = pagination.AddPagination(query)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, fmt.Errorf("Scan: %w", err)
	}

	return callList, nil
}
func (r *Repo) Update(ctx context.Context, call *call_entity.Call) error {
	result, err := r.db.NewUpdate().
		Model(call).
		Column("caller_id", "agent_id", "call_start", "call_end", "status_id", "call_notes").
		Where("id = ?", call.ID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("NewUpdate: %w", err)
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return &app_errors.CallNotFound{}
	}

	return nil
}
func (r *Repo) Delete(ctx context.Context, id int) error {
	result, err := r.db.NewDelete().
		Model((*call_entity.Call)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("NewDelete: %w", err)
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return &app_errors.CallNotFound{}
	}

	return nil
}
