package main

import (
	"fmt"
	"tic-tac-go/board"
	"tic-tac-go/player"
	"tic-tac-go/utils"
)

func main() {
	// init game conditions
	players := player.Init()
	b := board.Init()

	utils.DrawStartingScreen(b)
	
	// run the game loop
	for {
		currentPlayer := player.GetCurrentPlayer(players)
		x, y := utils.GetPlayerMove(b, currentPlayer)
		b.Place(x, y, currentPlayer.Token)
	
		b.Draw()
	
		if utils.IsGameOver(b, players) {
			fmt.Println(currentPlayer.Name, "wins!")
			break
		}
		if utils.IsDraw(b) {
			fmt.Println("It's a draw!")
			break
		}
	
		players = player.SwitchPlayersTurn(players)
	}
}
