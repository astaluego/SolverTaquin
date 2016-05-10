package main

import (
	"flag"
	"fmt"

	"github.com/asta-luego/n-puzzle/pkg"
)

func main() {
	var t taquin.Taquin
	flag.BoolVar(&t.Verbose, "v", false, "mode verbose")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		fmt.Println("  [file ...]")
	} else {
		if t.Verbose == true {
			fmt.Println("INITIALISATION")
			fmt.Println("Mode verbose : ", t.Verbose)
			fmt.Println("Filename : ", flag.Args())
		}
		t.Keepfiles(flag.Args())
	}

}
