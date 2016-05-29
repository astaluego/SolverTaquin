package taquin

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/apex/log"
)

// Taquin structure of a map
type Taquin struct {
	Filename string
	Size     int
	Board    [][]int
}

// Keepfiles check if file exists and add to the struct
func (t *Taquin) Keepfiles(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Couldn't open file %s", filename)
	}

	defer file.Close()

	nblines := 0
	line := 1
	var rline *regexp.Regexp
	var regexline string
	rcomment := regexp.MustCompile("^[#]+")
	rsize := regexp.MustCompile("^[-+]?[0-9]+$")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !rcomment.MatchString(scanner.Text()) {

			if rsize.MatchString(scanner.Text()) {
				// Check if size has not already been assigned
				if t.Size != 0 {
					log.Errorf("Invalid map in %s : the size is declared two times", filename)
					return
				}
				// Convert size to integer
				t.Size, err = strconv.Atoi(scanner.Text())
				if err != nil {
					log.Fatalf("Can't convert string to integer in %s : %v", filename, err)
				}
				if t.Size <= 0 {
					log.Errorf("Invalid map in %s : the size of map is %d", filename, t.Size)
					return
				}
				// Keep regex for lines
				for i := 0; i < t.Size; i++ {
					regexline = regexline + `[\s]*([0-9]+)`
				}
				rline = regexp.MustCompile(`^` + regexline + `[\s]*[#]*`)
				t.Board = make([][]int, t.Size)
			} else if t.Size >= 0 && rline != nil && rline.MatchString(scanner.Text()) {
				if nblines >= t.Size {
					log.Errorf("Invalid map in %s : too much lines", filename)
					return
				}

				res := rline.FindAllStringSubmatch(scanner.Text(), -1)

				elements := res[0][1:]
				t.Board[nblines] = make([]int, t.Size)
				for i, ele := range elements {
					t.Board[nblines][i], _ = strconv.Atoi(ele)
				}
				nblines++
			} else {
				log.Errorf("Invalid map in %s : error in line %d", filename, line)
				return
			}
			line++
		}
		if err = scanner.Err(); err != nil {
			log.Errorf("Error : %v", err)
			return
		}
	}
	fmt.Println(t.Board)
}

// CheckTaquin checks if is a valid map
func (t *Taquin) CheckTaquin() error {
	arrayCheck := make([]bool, (t.Size * t.Size))
	for i := range t.Board {
		for j := range t.Board[i] {
			if t.Board[i][j] > (t.Size*t.Size)-1 {
				return fmt.Errorf("Invalid map in %v : %v > %v", t.Filename, t.Board[i][j], (t.Size*t.Size)-1)
			}
			arrayCheck[t.Board[i][j]] = true
		}
	}
	for i, value := range arrayCheck {
		if !value {
			return fmt.Errorf("Inavlid map in %v : missing value %v", t.Filename, i)
		}
	}
	return nil
}
