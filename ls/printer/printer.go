package printer

import (
	"os"
)

type Printer interface {
	Print(os.File)
}
