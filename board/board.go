package board

import (
	"errors"
	"fmt"
)


type Board [3][3]string

func Init() Board {
	return Board{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

func (b Board) Draw() {
	fmt.Println()
	boardGrid := [5][5]string {
		{" ", "|", " ", "|", " "},
		{"-", "+", "-", "+", "-"},
		{" ", "|", " ", "|", " "},
		{"-", "+", "-", "+", "-"},
		{" ", "|", " ", "|", " "},
	}

	for x := range b {
		for y := range b[x] {
			val := b.Get(x, y)
			if val == "" {
				val = " "
			}
			boardGrid[x*2][y*2] = val
		}
	}

	for _, row:= range boardGrid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b Board) GetRaw(x, y int) string { // for game logic
	return b[x][y]
}

func (b Board) Get(x, y int) string { // for drawing
	if b[x][y] == "" {
		return " "
	}
	return b[x][y]
}

func (b *Board) Place(x, y int, token string) error {
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return errors.New("move out of bounds")
	}

	if b[x][y] != "" {
		return errors.New("cell already occupied")
	}

	b[x][y] = token
	return nil
}
