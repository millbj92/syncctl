// package disk

// import (
// 	"time"

// 	"io/fs"

// 	"github.com/google/uuid"
// )

// type DirectoryExt struct {
// 	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
// 	CreatedAt time.Time `db:"created_at" json:"created_at"`
// 	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

// 	Name               string  `db:"name" json:"name" validate:"required,lte=255"`
// 	Path               string  `db:"path" json:"path" validate:"required,lte=255"`
// 	Size               int64   `db:"size" json:"size"`
// 	Free               int64   `db:"free" json:"free"`
// 	Used               int64   `db:"used" json:"used"`
// 	Percentage         float64 `db:"percentage" json:"percentage"`
// 	Marked_For_Cleanup bool    `db:"marked_for_cleanup" json:"marked_for_cleanup"`

// 	base *fs.DirEntry // base directory
// }

// func (dr *DirectoryExt) New(base *fs.DirEntry) *DirectoryExt {
// 	d := &DirectoryExt{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		base:      base,
// 	}
// 	return d
// }
package disk
