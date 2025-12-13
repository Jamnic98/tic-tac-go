package player

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Player struct {
	Name string
	Token string
	IsCurrentPlayer bool
}

func Init() [2]Player{
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
		
		players[i] = Player{Name: playerNames[i], Token: playerIcon, IsCurrentPlayer: i == startingPlayerIndex}
	}

	return players
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

