package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Email string `db:"email" json:"email" validate:"required,email,lte=255"`
	Password string `db:"password" json:"password,omitempty" validate:"required,lte=255"`
	FirstName string `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName string `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Role string `db:"role" json:"role" validate:"required,lte=255" default:"disabled"`
}
