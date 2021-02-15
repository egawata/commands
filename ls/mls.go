package main

import (
	"flag"
	"log"

	"github.com/egawata/commands/ls/printer"
)

var (
	longFormat = flag.Bool("l", false, "Long format")
	withHidden = flag.Bool("a", false, "With hidden files")
)

func main() {
	var err error
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	var isMultiPath bool
	if len(paths) > 1 {
		isMultiPath = true
	}

	opt := &printer.PrinterOption{
		WithHidden: *withHidden,
		AddDirname: isMultiPath,
	}
	for _, path := range paths {
		var p printer.Printer

		if *longFormat {
			p = printer.NewLongPrinter(opt)
		} else {
			p = printer.NewSimplePrinter(opt)
		}

		err = p.Print(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
