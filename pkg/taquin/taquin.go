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
	Filename  string
	Size      int
	Board     [][]int
	GoalBoard [][]int
	EmptyCase []int
}

// Parse check if file exists and add to the struct
func (t *Taquin) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Couldn't open file %v", filename)
	}
	defer file.Close()

	t.Filename = filename
	fmt.Println("\nFILE --->", filename)
	var rline *regexp.Regexp

	nblines := 0
	j := 0
	rcomment := regexp.MustCompile("^[#]+")
	rsize := regexp.MustCompile("^[-+]?[0-9]+$")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return fmt.Errorf("Error : %v", err)
		}
		fmt.Println(scanner.Text())

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
				var regexline string
				for i := 0; i < t.Size; i++ {
					regexline = regexline + `[\s]*([0-9]+)`
				}
				rline = regexp.MustCompile(`^` + regexline + `(?:[\s]*[#]+[\s\S]*)?$`)
				t.Board = make([][]int, t.Size)
			} else if t.Size > 0 && rline != nil && rline.MatchString(scanner.Text()) {
				if nblines >= t.Size {
					return fmt.Errorf("Too much lines")
				}
				res := rline.FindAllStringSubmatch(scanner.Text(), -1)
				elements := res[0][1:]
				// Fill the map
				t.Board[nblines] = make([]int, t.Size)
				for i, ele := range elements {
					tmp, _ := strconv.Atoi(ele)
					t.Board[nblines][i] = tmp
					if tmp == 0 {
						t.EmptyCase = make([]int, 2)
						t.EmptyCase[0] = nblines
						t.EmptyCase[1] = i
					}
					j++
				}
				nblines++
			} else {
				return fmt.Errorf("Error invalid line")
			}
		}
	}
	if t.Size == 0 {
		return fmt.Errorf("You must specified a fucking size")
	}
	if t.Size != nblines {
		return fmt.Errorf("Missing line(s) in the map")
	}
	//fmt.Println(t.EmptyCase)
	t.MakeGoalBoard()
	t.PrintBoard(t.GoalBoard)
	t.PrintBoard(t.Board)
	return nil
}

// MakeGoalBoard Goal
func (t *Taquin) MakeGoalBoard() {
	t.GoalBoard = make([][]int, t.Size)
	for i := range t.GoalBoard {
		t.GoalBoard[i] = make([]int, t.Size)
		for j := range t.GoalBoard[i] {
			t.GoalBoard[i][j] = 0
		}
	}
	var dir = 0
	var value = 1
	var i = 0
	var j = 0
	for value < t.Size*t.Size {
		if value == t.Size*t.Size {
			t.GoalBoard[i][j] = 0
			break
		}
		if t.GoalBoard[i][j] == 0 {
			t.GoalBoard[i][j] = value
			value++
		}
		if dir == 0 {
			if j == t.Size-1 || t.GoalBoard[i][j+1] != 0 {
				dir = 1
				i++
			} else {
				j++
			}
		} else if dir == 1 {
			if i == t.Size-1 || t.GoalBoard[i+1][j] != 0 {
				dir = 2
				j--
			} else {
				i++
			}
		} else if dir == 2 {
			if j == 0 || t.GoalBoard[i][j-1] != 0 {
				dir = 3
				i--
			} else {
				j--
			}
		} else {
			if i == 0 || t.GoalBoard[i-1][j] != 0 {
				dir = 0
				j++
			} else {
				i--
			}
		}
	}
}

// CheckErrorParsingTaquin checks if is a valid map
func (t *Taquin) CheckErrorParsingTaquin() error {
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

// TranspositionInLineTaquin make un simple array of the board
func (t *Taquin) TranspositionInLineTaquin(board [][]int) []int {
	var k = 0
	var boardInLine = make([]int, t.Size*t.Size)
	for i := range board {
		for j := range board[i] {
			boardInLine[k] = board[i][j]
			k++
		}
	}
	return boardInLine
}

// Inversions calculate the number of inversions
func (t *Taquin) Inversions(board []int) (int, int) {
	var i = 0
	var posZero = 0
	var numberInversions = 0
	for i < t.Size*t.Size-1 {
		if board[i] == 0 {
			posZero = i
			i++
			continue
		}
		var j = i + 1
		for j < t.Size*t.Size {
			if board[j] == 0 {
				j++
				continue
			}
			if board[j] < board[i] {
				numberInversions++
			}
			j++
		}
		i++
	}
	return numberInversions, posZero
}

// CheckValidityTaquin checks if taquin is solvable
func (t *Taquin) CheckValidityTaquin() error {
	var boardInLine = t.TranspositionInLineTaquin(t.Board)
	var goalInLine = t.TranspositionInLineTaquin(t.GoalBoard)
	var goalInversions, goalZero = t.Inversions(goalInLine)
	var boardInversions, boardZero = t.Inversions(boardInLine)
	if t.Size%2 == 0 {
		goalInversions += goalZero / t.Size
		boardInversions += boardZero / t.Size
	}
	if goalInversions%2 != boardInversions%2 {
		return fmt.Errorf("Is not solvable")
	}
	return nil
}

// PrintBoard print board
//TODO
func (t *Taquin) PrintBoard(board [][]int) {
	var size = len(strconv.Itoa(t.Size * t.Size))
	fmt.Print("\n")
	for line := range board {
		for column := range board[line] {
			fmt.Printf("%*d", (size+1)*(-1), board[line][column])
		}
		fmt.Print("\n")
	}
}
