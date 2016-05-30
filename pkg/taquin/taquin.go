package taquin

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Taquin structure of a map
type Taquin struct {
	Filename string
	Size     int
	Board    [][]int
}

// Parse check if file exists and add to the struct
func (t *Taquin) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Couldn't open file %v", filename)
	}
	defer file.Close()

	t.Filename = filename

	var rline *regexp.Regexp
	var regexline string

	nblines := 0
	line := 0
	rcomment := regexp.MustCompile("^[#]+")
	rsize := regexp.MustCompile("^[-+]?[0-9]+$")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return fmt.Errorf("Error : %v", err)
		}
		line++
		if !rcomment.MatchString(scanner.Text()) {
			if rsize.MatchString(scanner.Text()) {
				// Check if size has not already been assigned
				if t.Size != 0 {
					return fmt.Errorf("The size is declared two times")
				}
				// Convert size to integer
				t.Size, err = strconv.Atoi(scanner.Text())
				if err != nil {
					return fmt.Errorf("Can't convert string to integer %v", err)
				}
				if t.Size <= 0 {
					return fmt.Errorf("The size of map is %v", t.Size)
				}
				// Keep regex for lines
				for i := 0; i < t.Size; i++ {
					regexline = regexline + `[\s]*([0-9]+)`
				}
				rline = regexp.MustCompile(`^` + regexline + `(?:[\s]*[#]+[\s\S]*)?$`)
				t.Board = make([][]int, t.Size)
			} else if t.Size >= 0 && rline != nil && rline.MatchString(scanner.Text()) {
				if nblines >= t.Size {
					return fmt.Errorf("Too much lines")
				}
				res := rline.FindAllStringSubmatch(scanner.Text(), -1)
				//fmt.Println(res)
				elements := res[0][1:]
				//Fill the map
				t.Board[nblines] = make([]int, t.Size)
				for i, ele := range elements {
					t.Board[nblines][i], _ = strconv.Atoi(ele)
				}
				nblines++
			} else {
				return fmt.Errorf("Error in line %v", line)
			}
		}
	}
	if t.Size == 0 {
		return fmt.Errorf("You must specified a fucking size")
	}
	if t.Size != nblines {
		return fmt.Errorf("Missing line(s) in the map")
	}
	fmt.Println(t.Board)
	return nil
}

// CheckTaquin checks if is a valid map
func (t *Taquin) CheckTaquin() error {
	arrayCheck := make([]bool, (t.Size * t.Size))
	for i := range t.Board {
		for j := range t.Board[i] {
			if t.Board[i][j] > (t.Size*t.Size)-1 {
				return fmt.Errorf("%v > %v", t.Board[i][j], (t.Size*t.Size)-1)
			}
			arrayCheck[t.Board[i][j]] = true
		}
	}
	for i, value := range arrayCheck {
		if !value {
			return fmt.Errorf("Missing value %v", i)
		}
	}
	return nil
}
