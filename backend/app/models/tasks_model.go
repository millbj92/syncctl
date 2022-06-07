package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	Name string `db:"name" json:"name" validate:"required,lte=255"`
	Interval string `db:"interval" json:"interval" validate:"required,lte=255"`
	Type string `db:"type" json:"type" validate:"required,lte=255"`
	Status string `db:"status" json:"status" validate:"required,lte=255"`
	LastRun time.Time `db:"last_run" json:"last_run"`
	NextRun time.Time `db:"next_run" json:"next_run"`

	StartAt time.Time `db:"start_at" json:"start_at"`
	EndAt time.Time `db:"end_at" json:"end_at"`


	Disabled bool `db:"disabled" json:"disabled"`
	Deleted bool `db:"deleted" json:"deleted"`
}
