package connections

import (
	"time"

	"github.com/google/uuid"
	"github.com/millbj92/synctl/libs/models/auth"
	"github.com/millbj92/synctl/libs/models/disk"
)

type ConnectionSettings struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	SSH_Key string `json:"ssh_key"`
	FriendlyName string `json:"friendly_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Poll_interval int `json:"poll_interval"`
	Poll_interval_unit string `json:"poll_interval_unit"`

	Sync_Directories []disk.DirectoryExt `json:"sync_directories"`
	Scan_Directories []disk.DirectoryExt `json:"scan_directories"`
	Move_To_Directories []disk.DirectoryExt `json:"move_to_directories"`
	Ignore_Directories []disk.DirectoryExt `json:"ignore_directories"`
	Ignore_Files []disk.FileExt `json:"ignore_files"`
	Include_Files []disk.FileExt `json:"include_files"`
	Exclude_Regex []string `json:"exclude_files_regex"`
	Sync_On_Startup bool `json:"sync_on_startup"`

	//ManagedProcesses []Process `json:"managed_processes"`
	Ignored_PIDs []int `json:"ignored_pids"`
	Ignored_Processes []string `json:"ignored_processes"`
	Ignore_Process_Pattern string `json:"ignore_process_pattern"`

	Ignore_CPU_Usage bool `json:"ignore_cpu_usage"`
	Ignore_Memory_Usage bool `json:"ignore_memory_usage"`
	Ignore_Disk_Usage bool `json:"ignore_disk_usage"`
	Ignore_Uptime bool `json:"ignore_uptime"`
	Ignore_Last_Seen bool `json:"ignore_last_seen"`
	Ignore_Processes bool `json:"ignore_processes"`
	Ignore_Process_Count bool `json:"ignore_process_count"`
	Ignore_Process_Memory bool `json:"ignore_process_memory"`
	Ignore_Process_CPU bool `json:"ignore_process_cpu"`
	Ignore_Process_Disk bool `json:"ignore_process_disk"`
	Ignore_Process_Uptime bool `json:"ignore_process_uptime"`

	Hidden bool `json:"hidden"`
	Login_Users []auth.User `json:"login_users"`
}
