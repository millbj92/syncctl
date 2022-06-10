package disk

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/google/uuid"
)

type FileExt struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Version int `db:"version" json:"version"`
	Marked_For_Deletion bool `db:"marked_for_deletion" json:"marked_for_deletion"`
	Marked_For_Cleanup bool `db:"marked_for_cleanup" json:"marked_for_cleanup"`
	Marked_For_Sync bool `db:"marked_for_sync" json:"marked_for_sync"`

	Name string `db:"name" json:"name" validate:"required,lte=255"`
	Path string `db:"path" json:"path" validate:"required,lte=255"`
	Size int64 `db:"size" json:"size"`

	Deleted_At time.Time `db:"deleted_at" json:"deleted_at"`

	isDeleted bool // used for soft delete
	base *fs.DirEntry // base directory

}

func (fl *FileExt) New(base *fs.DirEntry) *FileExt {
	f := &FileExt{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version: 1,
		Marked_For_Deletion: false,
		Marked_For_Cleanup: false,
		Marked_For_Sync: false,
		base: base,
	}
	return f
}

func (f *FileExt) GetBase() fs.DirEntry {
	return *f.base
}
func (f *FileExt) IsDeleted() bool {
	return f.isDeleted
}

func (f *FileExt) Delete() bool {
	f.Marked_For_Deletion = true
	return true
}

func (f *FileExt) SetDeleted(isDeleted bool) {
	f.isDeleted = isDeleted
}
func _Delete_From_Disk(f *FileExt) bool {
	f.Deleted_At = time.Now()
	err := os.Remove(f.Path); if err != nil {
		fmt.Println("Error deleting file: ", err)
		return false
	}
	return true;
}
