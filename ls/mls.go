package main

import (
	"flag"
	"log"
	"os"

	"github.com/egawata/commands/ls/printer"
)

var (
	withHidden = flag.Bool("all", false, "With hidden files")
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

	p := printer.NewLongPrinter(*withHidden)
	err = p.Print(f)
	if err != nil {
		log.Fatal(err)
	}
}
