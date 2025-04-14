package main

import (
	"fmt"
	"tic-tac-go/utils"
)

func main() {
	// init game conditions
	players := utils.InitPlayers()
	board := utils.InitBoard()

	utils.DrawStartingScreen(board)
	
	// run the game loop
	for {
		currentPlayer := utils.GetCurrentPlayer(players)
		x, y := utils.GetPlayerMove(board, currentPlayer)
		board[x][y] = currentPlayer.Token
	
		utils.DrawBoard(board)
	
		if utils.IsGameOver(board, players) {
			fmt.Println(currentPlayer.Name, "wins!")
			break
		}
		if utils.IsDraw(board) {
			fmt.Println("It's a draw!")
			break
		}
	
		players = utils.SwitchPlayersTurn(players)
	}
}
