package call_entity

import (
	"time"
)

type Status struct {
	ID    int `bun:"id,pk,autoincrement"`
	Title string
}

type Agent struct {
	ID    int `bun:"id,pk,autoincrement"`
	Login string
	FIO   string
}

type Call struct {
	ID        int `bun:"id,pk,autoincrement"`
	CallerID  string
	AgentID   int
	Agent     Agent `bun:"rel:belongs-to,join:agent_id=id"`
	CallStart time.Time
	CallEnd   time.Time
	StatusID  int
	Status    Status `bun:"rel:belongs-to,join:status_id=id"`
	CallNotes string `bun:",nullzero"`
}
