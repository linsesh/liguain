package services

import (
	"liguain/backend/models"
	"liguain/backend/rules"
)

type GameRepository interface {
	// GetGame returns a game if it exists
	GetGame(gameId string) (rules.Game, error)
	// SaveGame saves a game and returns the game id, and an error if saving failed
	SaveGame(game rules.Game) (string, error)
	updateScores(scores map[models.Player]int) error
}
