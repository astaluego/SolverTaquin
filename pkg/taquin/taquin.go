package taquin

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//Board this is the board game
type Board struct {
	Filename string
	Size     int
}

//Taquin taquin
type Taquin struct {
	Verbose bool
	Board
	//slice []Board
}

//Keepfiles check if file exists and add to the struct
func (t *Taquin) Keepfiles(filenames []string) {
	for i := range filenames {
		if file, err := os.Open(filenames[i]); err == nil {
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				rcomment, _ := regexp.Compile("^[#]+")
				if rcomment.MatchString(scanner.Text()) == false {
					rsize, _ := regexp.Compile("^[0-9]+$")
					if rsize.MatchString(scanner.Text()) == true {
						t.Size, _ = strconv.Atoi(scanner.Text())
					}
					// if t.Size != 0 {
					// 	rmap, _ := regexp.Compile("^(([0-9])+[ ]*){" + t.Size + "}$")
					// }
				}
			}
			if err = scanner.Err(); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(err)
			return
		}
		t.Filename = filenames[i]
	}
	fmt.Println(t.Filename)
	fmt.Println(t.Size)
}
