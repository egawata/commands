package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/egawata/commands/ls/printer"
)

var (
	longFormat = flag.Bool("l", false, "Long format")
	withHidden = flag.Bool("a", false, "With hidden files")
)

func main() {
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	var isMultiPath bool
	if len(paths) > 1 {
		isMultiPath = true
	}

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		var p printer.Printer

		if *longFormat {
			p = printer.NewLongPrinter(*withHidden)
		} else {
			p = printer.NewSimplePrinter(*withHidden)
		}

		if isMultiPath {
			fmt.Printf("%s:\n", path)
		}
		err = p.Print(f)
		if err != nil {
			log.Fatal(err)
		}
		if isMultiPath {
			fmt.Printf("\n")
		}
	}
}
