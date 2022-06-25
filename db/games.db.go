package db

import (
	"fmt"

	"github.com/iain-apw/wordle_game/helpers"
	"github.com/iain-apw/wordle_game/models"
)

type GameDB struct {
	games []models.Game
}

// NewGames creates a new empty games db
func NewGames() (*GameDB, error) {
	var games = []models.Game{
		helpers.NewGame(5, &models.User{ID: "12345", Name: "Numbers"}),
		helpers.NewGame(6, &models.User{ID: "ABCDE", Name: "Letters"}),
		helpers.NewGame(5, &models.User{ID: "Jeff", Name: "Jeffrey"}),
		helpers.NewGame(9, &models.User{ID: "Mitch", Name: "David Mitchell"}),
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

// Find all games by user
func (g *GameDB) FindAllByUser(userID string) ([]models.Game, error) {
	userGames := []models.Game{}

	for _, game := range g.games {
		if game.UserId == userID {
			userGames = append(userGames, game)
		}
	}

	return userGames, nil
}

// AddGame Adds a new game to the system
func (g *GameDB) AddGame(newGame models.Game) (models.Game, error) {
	//var allGames []models.Game
	g.games = append(g.games, newGame)

	return newGame, nil
}

func (g *GameDB) UpdateGame(updatedGame models.Game) (models.Game, error) {
	updated := false

	for i := 0; i < len(g.games); i++ {

		if g.games[i].ID == updatedGame.ID {
			g.games[i] = updatedGame
			updated = true
			break
		}
	}

	if !updated {
		return models.Game{}, fmt.Errorf("game id %s not found", updatedGame.ID)
	}

	return updatedGame, nil
}
