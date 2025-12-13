package utils

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"tic-tac-go/board"
	"tic-tac-go/player"
)


func IsGameOver(b board.Board, players [2]player.Player) bool {
	magicSquare := [3][3]int{
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8},
	}

	// all possible winning combinations in a 3x3 magic square
	winningCombos := [][]int{
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8},
		{2, 9, 4},
		{7, 5, 3},
		{6, 1, 8},
		{2, 5, 8},
		{6, 5, 4},
	}

	for _, player := range players {
		var playerMoves []int

		for x := range b {
			for y := range b[x] {
				if b.Get(x, y) == player.Token {
					playerMoves = append(playerMoves, magicSquare[x][y])
				}
			}
		}

		// Check if player has any of the winning combinations
		for _, combo := range winningCombos {
			if hasAll(playerMoves, combo) {
				return true
			}
		}
	}
	return false
}

func hasAll(moves []int, combo []int) bool {
	count := 0
	for _, val := range combo {
		if slices.Contains(moves, val) {
				count++
			}
	}
	return count == 3
}

func IsDraw(b board.Board) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if b.GetRaw(x, y) == "" {
				return false
			}
		}
	}
	return true
}


func GetPlayerMove(b board.Board, currentPlayer player.Player) (int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(currentPlayer.Name, "'s turn. Enter move in format x,y: ")
		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			continue
		}

		move := scanner.Text()
		coords := strings.Split(move, ",")
		if len(coords) != 2 {
			fmt.Println("Invalid input format. Please enter in x,y format.")
			continue
		}

		x64, err := strconv.ParseInt(strings.TrimSpace(coords[0]), 10, 0)
		if err != nil || x64 < 0 || x64 > 2 {
			fmt.Println("Invalid x coordinate. Must be 0, 1, or 2.")
			continue
		}

		y64, err := strconv.ParseInt(strings.TrimSpace(coords[1]), 10, 0)
		if err != nil || y64 < 0 || y64 > 2 {
			fmt.Println("Invalid y coordinate. Must be 0, 1, or 2.")
			continue
		}

		x, y := int(x64), int(y64)
		row := 2 - y    // flip y to match row index
		col := x        // x is column


		if b.GetRaw(row, col) != "" {
			fmt.Println("Cell is not empty, choose another one!")
			b.Draw()
			continue
		}

		return row, col
	}
}

func DrawStartingScreen(b board.Board) {
	fmt.Println()
	fmt.Println("Game starting!")
	b.Draw()
}
