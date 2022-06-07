package disk

import (
	"time"
)

type Drive struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Free int64 `json:"free"`
	Used int64 `json:"used"`
	Percentage float64 `json:"percentage"`
	Volumes []Volume `json:"volumes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Drive) New() *Drive {

	return &Drive{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
