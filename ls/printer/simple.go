package printer

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// ファイルを横に並べるときの間隔
	marginX   = 3
	widthIcon = len(FILE_ICON_NORMAL) + 1
)

type SimplePrinter struct {
	withHidden bool
}

func NewSimplePrinter(withHidden bool) *SimplePrinter {
	return &SimplePrinter{
		withHidden: withHidden,
	}
}

func (p *SimplePrinter) Print(f *os.File) error {
	pi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Stat: %w", err)
	}

	termWidth, _, err := terminal.GetSize(syscall.Stdout)
	if err != nil {
		return fmt.Errorf("GetSize: %w", err)
	}

	if pi.IsDir() {
		files, err := f.ReadDir(0)
		if err != nil {
			return fmt.Errorf("ReadDir: %w", err)
		}

		longest := 0
		var iList []os.FileInfo

		for _, f := range files {
			i, err := f.Info()
			if err != nil {
				return fmt.Errorf("Info: %w", err)
			}

			name := i.Name()
			if !p.withHidden {
				if name[0] == '.' {
					continue
				}
			}

			if longest < len(name) {
				longest = len(name)
				fmt.Printf("longest = %s, %d\n", name, len(name))
			}

			iList = append(iList, i)
		}
		if longest == 0 {
			return nil
		}

		colWidth := longest + marginX + widthIcon
		colNum := int(termWidth / colWidth)
		if colNum == 0 {
			colNum = 1
		}
		rowNum := int(len(iList)/colNum) + 1
		fmt.Printf("termwidth = %d, col = %d, row = %d, longest = %d\n", termWidth, colNum, rowNum, longest)

		for y := 0; y < rowNum; y++ {
			for x := 0; x < colNum; x++ {
				ind := x*rowNum + y
				if ind >= len(iList) {
					continue
				}
				p.printFile(iList[ind], colWidth)
			}
			fmt.Printf("\n")
		}
	} else {
		p.printFile(pi, termWidth)
	}

	return nil
}

func (p *SimplePrinter) printFile(i os.FileInfo, colWidth int) {
	var filePrefix = FILE_ICON_NORMAL
	if i.IsDir() {
		color.Set(color.FgBlue)
		filePrefix = FILE_ICON_DIR
	}
	printed := fmt.Sprintf("%s %s", filePrefix, i.Name())
	if colWidth-len(printed) < 0 {
		log.Fatalf("colWidth = %d, len = %d, string = %s", colWidth, len(printed), printed)
	}
	fmt.Printf("%s%s", printed, strings.Repeat(" ", colWidth-len(printed)))
	color.Unset()
}
