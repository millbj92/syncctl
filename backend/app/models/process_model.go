package models

import "time"

type Process struct {
	Pid int `json:"pid"`
	Name string `json:"name"`
	User string `json:"user"`
	State string `json:"state"`
	Memory int64 `json:"memory"`
	CPU float64 `json:"cpu"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
