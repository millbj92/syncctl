package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/millbj92/synctl/app/models/disk"
)

type Connection struct {
	Id uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	IP_Address string `json:"ip_address"`
	Hostname string `json:"hostname"`
	FriendlyName string `json:"friendly_name"`
	Port int `json:"port"`
	Protocol string `json:"protocol"`
	Username string `json:"username"`
	SSH_Key string `json:"ssh_key"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Settings ConnectionSettings `json:"settings"`

	Memory_Total int `json:"memory_total"`
	Memory_Free int `json:"memory_free"`
	Top_ten_processes []Process `json:"top_ten_processes"`

	CPU_Total int `json:"cpu_total"`
	CPU_Free int `json:"cpu_free"`
	CPU_Cores int `json:"cpu_cores"`

	Drives []disk.Drive `json:"drives"`
	Swap_Total int `json:"swap_total"`
	Swap_Free int `json:"swap_free"`

	Uptime int `json:"uptime"`

	Last_seen time.Time `json:"last_seen"`


}
