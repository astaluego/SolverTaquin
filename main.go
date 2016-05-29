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
		log.Error("Please give me a filename")
	} else {
		if verbose {
			os.Setenv("TAQUIN_VERBOSE", "1")
		}
		t := make([]taquin.Taquin, nbfile)
		for i, filename := range flag.Args() {
			t[i].Keepfiles(filename)
			if err := t[i].CheckTaquin(); err != nil {
				log.Errorf("%v", err)
			}
		}
	}
}
