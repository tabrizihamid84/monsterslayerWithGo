package main

import (
	"github.com/tabrizihamid84/monsterSlayer/interaction"
)

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	return ""
}

func endGame() {}
