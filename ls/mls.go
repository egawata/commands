package main

import (
	"flag"
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

	args := flag.Args()
	var path string
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}

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

	err = p.Print(f)
	if err != nil {
		log.Fatal(err)
	}
}
