package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)


type Player struct {
	Name string
	Token string
	IsCurrentPlayer bool
}

func InitBoard() [3][3]string {
return [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

func InitPlayers() [2]Player{
scanner := bufio.NewScanner(os.Stdin)

icons := [2]string{"g", "o"}
var players [2]Player
var playerNames []string

startingPlayerIndex := rand.Intn(2)
for i := range players {
	playerIcon := icons[i]
	fmt.Print("Enter player ", playerIcon, "'s name: ")
	for scanner.Scan() {
		playerNames = append(playerNames, scanner.Text())
		break
	}
	
	players[i] = Player{playerNames[i], playerIcon, i == startingPlayerIndex}
}

return players
}

func DrawBoard(board [3][3]string) {
	fmt.Println()
	boardGrid := [5][5]string {
		{" ", "|", " ", "|", " "},
		{"-", "+", "-", "+", "-"},
		{" ", "|", " ", "|", " "},
		{"-", "+", "-", "+", "-"},
		{" ", "|", " ", "|", " "},
	}

	for x := range board {
		for y := range board[x] {
			val := board[x][y]
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

func IsGameOver(board [3][3]string, players [2]Player) bool {
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

		for x := range board {
			for y := range board[x] {
				if board[x][y] == player.Token {
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

func IsDraw(board [3][3]string) bool {
for x := range board {
	for y := range board[x] {
		if board[x][y] == "" {
			return false
		}
	}
}
return true
}


func GetPlayerMove(board [3][3]string, currentPlayer Player) (int, int) {
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

	if board[x][y] != "" {
		fmt.Println("Cell is not empty, choose another one!")
		DrawBoard(board)
		continue
	}

	return x, y
}
}

func GetCurrentPlayer(players [2]Player) Player {
	var currentPlayer Player 
	for _, player := range players {
		if player.IsCurrentPlayer {
			currentPlayer = player
		}
	}
	return currentPlayer
	}

func SwitchPlayersTurn(players [2]Player) [2]Player {
	players[0].IsCurrentPlayer = !players[0].IsCurrentPlayer
	players[1].IsCurrentPlayer = !players[1].IsCurrentPlayer
	return players
}

func DrawStartingScreen(board [3][3]string) {
	fmt.Println()
	fmt.Println("Game starting!")
	DrawBoard(board)
}
