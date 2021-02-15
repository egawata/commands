package printer

import (
	"os"
)

const (
	FILE_ICON_NORMAL = "\U0001f4c3"
	FILE_ICON_DIR    = "\U0001f4c1"
)

type Printer interface {
	Print(*os.File) error
}
