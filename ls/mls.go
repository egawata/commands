package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
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

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		i, err := f.Info()
		if err != nil {
			log.Fatal(err)
		}

		if !*withHidden {
			if i.Name()[0] == '.' {
				continue
			}
		}
		fmt.Printf("%10d\t%s\t", i.Size(), i.ModTime().Format("2006-01-02 15:04"))

		if i.IsDir() {
			color.Set(color.FgBlue)
		}
		fmt.Printf("%s", i.Name())
		color.Unset()

		fmt.Printf("\n")
	}
}
