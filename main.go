package main

import (
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/asta-luego/npuzzle/pkg/taquin"
)

var verbose bool

func init() {
	log.SetHandler(text.New(os.Stderr))
	flag.BoolVar(&verbose, "v", false, "mode verbose")
	flag.Parse()
}

func main() {
	nbfile := len(flag.Args())

	if nbfile == 0 {
		log.Fatal("Please give me a filename")
	}
	if verbose {
		os.Setenv("TAQUIN_VERBOSE", "1")
	}
	t := make([]taquin.Taquin, nbfile)
	for i, filename := range flag.Args() {
		if err := t[i].Parse(filename); err != nil {
			log.Errorf("Invalid map %v : %v", filename, err)
		} else if err = t[i].CheckErrorParsingTaquin(); err != nil {
			log.Errorf("Invalid map %v : %v", filename, err)
		} else if err = t[i].CheckValidityTaquin(); err != nil {
			log.Errorf("Invalid map %v : %v", filename, err)
		}
		//t[i].PrintBoard()
	}
}
