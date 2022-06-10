package scheduling

import (
	"github.com/google/uuid"
)

type Schedule struct {
	ID           uuid.UUID `db:"id" json:"id,omitempty" validate:"required,uuid"`
	Name         string    `db: "name" json:"name,omitempty" validate:"required"`
	Description  string    `db: "description" json:"description`
	TaskId       string    `db: "task_id" json:"task_id "validate: "required,uuid"`
	ConnectionId int64     `db: "connection_id" json:"connection_id validate: "required,uuid"`
	Enabled      bool      `db: "enabled" json:"enabled" validate: "required"`
	Cron         string    `db: "cron" json:"cron" validate: "required"`
}
