package printer

import (
	"os"
)

const (
	FILE_ICON_NORMAL = "\U0001f4c3"
	FILE_ICON_DIR    = "\U0001f4c1"
)

type Printer interface {
	Print(string) error
}

type PrinterOption struct {
	WithHidden bool
	AddDirname bool
}

type fileInfoList []os.FileInfo

func (f fileInfoList) Len() int {
	return len(f)
}

func (f fileInfoList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f fileInfoList) Less(i, j int) bool {
	return f[i].Name() < f[j].Name()
}
