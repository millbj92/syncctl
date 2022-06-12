package monitoring

import (
	//"io"
	//"fmt"
	"os"
	"path/filepath"

	"io/fs"

	"github.com/shirou/gopsutil/v3/disk"

	"github.com/davecgh/go-spew/spew"
	//github.com/variantdev/chartify
	//https://github.com/variantdev/vals
	//go get gopkg.in/yaml.v3
	//"github.com/gosuri/uitable"
	//"github.com/logrusorgru/aurora"
	//"github.com/sirupsen/logrus"
	//prettytable "github.com/tatsushid/go-prettytable"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type DiskOperation int64

const (
	NOOP DiskOperation = iota
	Read
	COPY
	MOVE
	RENAME
	DELETE
)

type DiskOpRequest struct {
	Operation   DiskOperation
	SourceDir   string
	SearchGLOB  string
	DestDir     string
	IgnoreTypes []string
	IgnoreDirs  []string
	IgnoreFiles []string
}

func (d DiskOpRequest) IsIgnored(path string) bool {
	if contains(d.IgnoreDirs, filepath.Base(path)) {
		return true
	}
	return false
}

func (d DiskOpRequest) IsIgnoredFile(path string) bool {
	if contains(d.IgnoreFiles, filepath.Base(path)) {
		return true
	}
	return false
}

func (d DiskOpRequest) IsIgnoredType(path string) bool {
	if contains(d.IgnoreTypes, filepath.Ext(path)) {
		return true
	}
	return false
}

func (d DiskOpRequest) IsValid() bool {
	if d.Operation == NOOP {
		return false
	}
	if d.SourceDir == "" {
		return false
	}
	if d.DestDir == "" {
		return false
	}
	return true
}

func (d DiskOpRequest) String() string {
	return spew.Sdump(d)
}

func (d DiskOpRequest) GetSourceDir() string {
	return d.SourceDir
}

// func (d DiskOpRequest) Execute() error {
// 	if d.Operation == NOOP {
// 		return nil
// 	}
// 	if d.Operation == COPY {
// 		return d.Copy()
// 	}
// 	if d.Operation == MOVE {
// 		return d.Move()
// 	}
// 	if d.Operation == RENAME {
// 		return d.Rename()
// 	}
// 	if d.Operation == DELETE {
// 		return d.Delete()
// 	}
// 	return nil
// }

func DiskUsage(path string) (*disk.UsageStat, error) {
	st, err := disk.Usage(path)
	if err != nil {
		spew.Dump(err)
	}

	return st, nil

	/*green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("Disk Total: %s\n", green(st.Total))
	fmt.Printf("Disk Free: %s\n", green(st.Free))
	fmt.Printf("Total Used: %s\n", red(st.Used))

	fmt.Printf("Used Percent: %s\n", red(pctStr))*/

}

func PrintDiskUsageByPath(path string) (uint64, error) {
	st, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}
	return st.Used, nil
}

func PrintPartitions() {
	st, err := disk.Partitions(true)
	if err != nil {
		spew.Dump(err)
	}

	headerFmt := color.New(color.FgGreen, color.Bold, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgCyan).SprintfFunc()

	tbl := table.New("Device", "Mountpoint", "FSType")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, v := range st {
		tbl.AddRow(v.Device, v.Mountpoint, v.Fstype)
	}
	tbl.Print()
}

func ScanDirs(path string) {
	var extensions = []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".tif", ".svg", ".webp", ".ico", ".psd", ".raw", ".arw", ".cr2", ".nef", ".dng", ".rw2", ".orf", ".raf", ".srw", ".x3f", ".rwl", ".crw", ".cr2", ".kdc", ".dcr", ".drf", ".dng", ".erf", ".fff", ".mef", ".mos", ".mrw", ".pef", ".ptx", ".raw", ".rwl", ".rw2", ".sr2", ".srf", ".srw", ".x3f", ".xmp", ".jpeg", ".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp", ".jfif-tbnl", ".jpe", ".jfif-tbnl", ".jfi", ".jfif-tiff", ".jfif-webp", ".jpm", ".jpgm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx", ".jp2", ".j2k", ".j2c", ".jpc", ".j2i", ".jfif-jpeg", ".jpm", ".jpx",}
	fsys := os.DirFS(path)
	err := fs.WalkDir(fsys, ".",
		func(p string,
			d fs.DirEntry, err error) error {
			i, err := d.Info();
			if err != nil {
				spew.Dump(err)
				return err
			}
			if contains(extensions, filepath.Ext(p)) {
				spew.Dump(i.Name())
			}
			return nil
		})
	if err != nil {
		return
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// fs.WalkDir(fsys, '.', func(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	dirs.Add(info, info.Size())
// 	return nil
// }, err error,

// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(.prettytable.?)
