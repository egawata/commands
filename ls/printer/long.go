package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

type LongPrinter struct {
	withHidden bool // 隠しファイルも表示する
	addDirname bool
}

func NewLongPrinter(opt *PrinterOption) *LongPrinter {
	return &LongPrinter{
		withHidden: opt.WithHidden,
		addDirname: opt.AddDirname,
	}
}

func (p *LongPrinter) Print(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open: %w", err)
	}
	defer f.Close()

	pi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Stat: %w", err)
	}

	if pi.IsDir() {
		files, err := f.ReadDir(0)
		if err != nil {
			return fmt.Errorf("ReadDir: %w", err)
		}

		var fis fileInfoList
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
			fis = append(fis, i)
		}
		sort.Sort(fis)

		if p.addDirname {
			fmt.Printf("%s:\n", path)
		}

		for _, fi := range fis {
			p.printFile(fi)
		}
	} else {
		matches, err := filepath.Glob(pi.Name())
		if err != nil {
			return fmt.Errorf("Glob: %w", err)
		}
		for _, fn := range matches {
			f, err := os.Open(fn)
			if err != nil {
				return fmt.Errorf("Open: %w", err)
			}
			i, err := f.Stat()
			if err != nil {
				return fmt.Errorf("Info: %w", err)
			}
			p.printFile(i)
		}
	}

	return nil
}

func (p *LongPrinter) printFile(i os.FileInfo) {
	fmt.Printf("%10d\t%s\t", i.Size(), i.ModTime().Format("2006-01-02 15:04"))

	var filePrefix = FILE_ICON_NORMAL
	if i.IsDir() {
		color.Set(color.FgBlue)
		filePrefix = FILE_ICON_DIR
	}
	fmt.Printf("%s", filePrefix+" "+i.Name())
	color.Unset()

	fmt.Printf("\n")
}
