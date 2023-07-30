package main

import (
	"github.com/tabrizihamid84/monsterSlayer/actions"
	"github.com/tabrizihamid84/monsterSlayer/interaction"
)

var currentRound = 0

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
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {
		actions.AttackMonster(isSpecialRound)

	} else if userChoice == "HEAL" {
		actions.HealPlayer()

	} else {
		actions.AttackMonster(true)
	}

	actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {}
