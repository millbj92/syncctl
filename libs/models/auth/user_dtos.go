package auth


type UserForCreate struct {
	Email string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	FirstName string `json:"first_name" validate:"required,lte=255"`
	LastName string `json:"last_name" validate:"required,lte=255"`
	Role string `json:"role" validate:"required,lte=255" default:"disabled"`
}

type UserForLogin struct {
	Email string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
