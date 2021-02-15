package printer

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type LongPrinter struct {
	withHidden bool // 隠しファイルも表示する
}

func NewLongPrinter(withHidden bool) *LongPrinter {
	return &LongPrinter{
		withHidden: withHidden,
	}
}

func (p *LongPrinter) Print(f *os.File) error {
	pi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Stat: %w", err)
	}

	if pi.IsDir() {
		files, err := f.ReadDir(0)
		if err != nil {
			return fmt.Errorf("ReadDir: %w", err)
		}

		for _, f := range files {
			i, err := f.Info()
			if err != nil {
				return fmt.Errorf("Info: %w", err)
			}

			if !p.withHidden {
				if i.Name()[0] == '.' {
					continue
				}
			}
			printFile(i)
		}
	} else {
		printFile(pi)
	}

	return nil
}

func printFile(i os.FileInfo) {
	fmt.Printf("%10d\t%s\t", i.Size(), i.ModTime().Format("2006-01-02 15:04"))

	var filePrefix = "\U0001f4c3"
	if i.IsDir() {
		color.Set(color.FgBlue)
		filePrefix = "\U0001f4c1"
	}
	fmt.Printf("%s", filePrefix+" "+i.Name())
	color.Unset()

	fmt.Printf("\n")
}
