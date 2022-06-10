package disk

import "time"

type Volume struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Free int64 `json:"free"`
	Used int64 `json:"used"`
	Percentage float64 `json:"percentage"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Mounted bool `json:"mounted"`
	MountPoint string `json:"mount_point"`
	Filesystem string `json:"filesystem"`
	Top_Ten_Directories []DirectoryExt `json:"top_ten_directories"`
}
