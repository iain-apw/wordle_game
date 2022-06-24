package repo

import (
	"fmt"

	"github.com/iain-apw/wordle_game/db"
	"github.com/iain-apw/wordle_game/models"
)

// repo holds all the dependencies required for repo operations
type gamesRepo struct {
	games *db.GameDB
	/*
		games
		products *db.ProductDB
		orders   *db.OrderDB
	*/
}

// Repo is the interface we expose to outside packages
type GamesRepo interface {
	GetAllGames() ([]models.Game, error)
	GetGame(id string) (models.Game, error)
	GetCurrentGame(user *models.User) (models.Game, error)
	CreateGame(request models.CreateGameRequest, user *models.User) (models.Game, error)
}

// New creates a new  GamesRepo with the correct database dependencies
func New() (GamesRepo, error) {
	g, err := db.NewGames()
	if err != nil {
		return nil, err
	}

	r := gamesRepo{
		games: g,
	}
	return &r, nil
}

// GetAllGames returns all games in the system
func (r *gamesRepo) GetAllGames() ([]models.Game, error) {
	return r.games.FindAll(), nil
}

// GetAllGames returns all games in the system
func (r *gamesRepo) GetGame(id string) (models.Game, error) {
	return r.games.FindById(id)
}

func (r *gamesRepo) GetCurrentGame(user *models.User) (models.Game, error) {
	allUserGames, err := r.games.FindAllByUser(user.ID)

	if err != nil {
		return models.Game{}, err
	}

	currentGame := models.Game{}
	found := false

	for _, game := range allUserGames {
		if game.Status == models.GameStatus_InProgress {
			currentGame = game
			found = true
			break
		}
	}

	if !found {
		return currentGame, fmt.Errorf("no current game found for user id %s", user.ID)
	}

	return currentGame, nil
}

// CreateGame creates a new game in the system
func (r *gamesRepo) CreateGame(request models.CreateGameRequest, user *models.User) (models.Game, error) {
	currentGame, _ := r.GetCurrentGame(user)

	if currentGame.ID != "" {
		return currentGame, fmt.Errorf("game %s already in progress", currentGame.ID)
	}

	// Generate a new word from the word lists

	newGame := models.NewGame(request.Letters, user)
	r.games.AddGame(newGame)
	return newGame, nil
}
