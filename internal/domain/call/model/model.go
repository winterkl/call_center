package call_model

import (
	"contact_center/internal/app_errors"
	call_entity "contact_center/internal/domain/call/entity"
	"time"
)

type CreateCallRequest struct {
	CallerID  string    `json:"caller_id" binding:"required"`
	AgentID   int       `json:"agent_id" binding:"required"`
	CallStart time.Time `json:"call_start" binding:"required"`
	CallEnd   time.Time `json:"call_end" binding:"required"`
	StatusID  int       `json:"status_id" binding:"required"`
	CallNotes string    `json:"call_notes"`
}

func (m *CreateCallRequest) GetEntity() *call_entity.Call {
	return &call_entity.Call{
		CallerID:  m.CallerID,
		AgentID:   m.AgentID,
		CallStart: m.CallStart,
		CallEnd:   m.CallEnd,
		StatusID:  m.StatusID,
		CallNotes: m.CallNotes,
	}
}

func (m *CreateCallRequest) Validate() error {
	if m.CallerID == "" {
		return &app_errors.IsRequired{Field: "caller_id"}
	}
	if m.AgentID == 0 {
		return &app_errors.IsRequired{Field: "agent_id"}
	}
	if m.CallStart.IsZero() {
		return &app_errors.IsRequired{Field: "call_start"}
	}
	if m.CallEnd.IsZero() {
		return &app_errors.IsRequired{Field: "call_end"}
	}
	if m.StatusID == 0 {
		return &app_errors.IsRequired{Field: "status_id"}
	}
	return nil
}

type CallFilter struct {
	CallerID   []string  `form:"caller_id"`
	AgentID    []int     `form:"agent_id"`
	CallStatus []int     `form:"status_id"`
	Begin      time.Time `form:"begin"`
	End        time.Time `form:"end"`
}

func (f *CallFilter) Validate() error {
	if f.Begin.IsZero() || f.End.IsZero() {
		return &app_errors.PeriodIsRequired{}
	}
	if f.Begin.After(f.End) {
		return &app_errors.BeginAfterEnd{}
	}
	return nil
}

type GetCallResponse struct {
	ID       int    `json:"id"`
	CallerID string `json:"caller_id"`
	Agent    struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		FIO   string `json:"fio"`
	} `json:"agent"`
	CallStart time.Time `json:"call_start"`
	CallEnd   time.Time `json:"call_end"`
	Status    struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"status"`
	CallNotes string `json:"call_notes"`
}

func NewGetCallResponse(call *call_entity.Call) *GetCallResponse {
	return &GetCallResponse{
		ID:       call.ID,
		CallerID: call.CallerID,
		Agent: struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			FIO   string `json:"fio"`
		}{
			ID:    call.Agent.ID,
			Login: call.Agent.Login,
			FIO:   call.Agent.FIO,
		},
		CallStart: call.CallStart,
		CallEnd:   call.CallEnd,
		Status: struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		}{
			ID:    call.StatusID,
			Title: call.Status.Title,
		},
		CallNotes: call.CallNotes,
	}
}

func GetCallResponseList(callList []call_entity.Call) []GetCallResponse {
	responseList := make([]GetCallResponse, len(callList))
	for ix, call := range callList {
		responseList[ix] = *NewGetCallResponse(&call)
	}
	return responseList
}

type UpdateCallRequest struct {
	ID        int
	CallerID  string    `json:"caller_id" binding:"required"`
	AgentID   int       `json:"agent_id" binding:"required"`
	CallStart time.Time `json:"call_start" binding:"required"`
	CallEnd   time.Time `json:"call_end" binding:"required"`
	StatusID  int       `json:"status_id" binding:"required"`
	CallNotes string    `json:"call_notes" binding:"required"`
}

func (m *UpdateCallRequest) Validate() error {
	if m.CallerID == "" {
		return &app_errors.IsRequired{Field: "caller_id"}
	}
	if m.AgentID == 0 {
		return &app_errors.IsRequired{Field: "agent_id"}
	}
	if m.CallStart.IsZero() {
		return &app_errors.IsRequired{Field: "call_start"}
	}
	if m.CallEnd.IsZero() {
		return &app_errors.IsRequired{Field: "call_end"}
	}
	if m.StatusID == 0 {
		return &app_errors.IsRequired{Field: "status_id"}
	}
	return nil
}

func (m *UpdateCallRequest) GetEntity() *call_entity.Call {
	return &call_entity.Call{
		ID:        m.ID,
		CallerID:  m.CallerID,
		AgentID:   m.AgentID,
		CallStart: m.CallStart,
		CallEnd:   m.CallEnd,
		StatusID:  m.StatusID,
		CallNotes: m.CallNotes,
	}
}
