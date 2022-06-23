package db

import (
	"fmt"

	"github.com/iain-apw/wordle_game/models"
)

type GameDB struct {
	games []models.Game
}

// NewGames creates a new empty games db
func NewGames() (*GameDB, error) {
	var games = []models.Game{
		models.NewGame(5, &models.User{ID: "12345", Name: "Numbers"}),
		models.NewGame(6, &models.User{ID: "ABCDE", Name: "Letters"}),
		models.NewGame(5, &models.User{ID: "Jeff", Name: "Jeffrey"}),
	}

	g := &GameDB{
		games: games,
	}

	return g, nil
}

// FindAll returns all games in the system
func (g *GameDB) FindAll() []models.Game {
	//var allGames []models.Game

	return g.games
}

// FindById returns a game in the system by it's id
func (g *GameDB) FindById(id string) (models.Game, error) {
	matchedGame := models.Game{}
	found := false

	for _, game := range g.games {
		if game.ID == id {
			matchedGame = game
			found = true
			break
		}
	}

	if !found {
		return matchedGame, fmt.Errorf("no game found for %s game id", id)
	}

	return matchedGame, nil
}

// AddGame Adds a new game to the system
func (g *GameDB) AddGame(newGame models.Game) (models.Game, error) {
	//var allGames []models.Game
	g.games = append(g.games, newGame)

	return newGame, nil
}
