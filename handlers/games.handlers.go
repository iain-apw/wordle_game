package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iain-apw/wordle_game/ctx"
	"github.com/iain-apw/wordle_game/helpers"
	"github.com/iain-apw/wordle_game/models"
	"github.com/iain-apw/wordle_game/repo"
)

type handler struct {
	repo repo.GamesRepo
}

type Handler interface {
	GetAllGames(w http.ResponseWriter, req *http.Request)
	GetGame(w http.ResponseWriter, req *http.Request)
	CreateGame(w http.ResponseWriter, req *http.Request)
	GetCurrentGame(w http.ResponseWriter, r *http.Request)
	MakeGuess(w http.ResponseWriter, r *http.Request)
}

func New() (Handler, error) {
	r, err := repo.New()
	if err != nil {
		return nil, err
	}
	h := handler{repo: r}
	return &h, nil
}

func (h *handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	allGames, err := h.repo.GetAllGames()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allGames)
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]

	game, err := h.repo.GetGame(gameId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func (h *handler) GetCurrentGame(w http.ResponseWriter, r *http.Request) {
	user, err := ctx.GetUserFromContext(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game, err := h.repo.GetCurrentGame(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

// OrderInsert creates a new order with the given parameters
func (h *handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var request models.CreateGameRequest

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctx.GetUserFromContext(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game, err := h.repo.CreateGame(request, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func (h *handler) MakeGuess(w http.ResponseWriter, r *http.Request) {
	var request models.MakeGuessRequest

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctx.GetUserFromContext(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get Current Game
	game, err := h.repo.GetCurrentGame(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Is word valid?
	isValid, err := helpers.IsValidWord(request.Guess, game.Letters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response = models.MakeGuessResponse{
		IsValidWord: isValid,
		Guess:       models.Guess{},
	}

	if isValid {
		response.Guess = helpers.CheckGuess(request.Guess, game.Answer)

		updatedGame, err := h.repo.UpdateGame(game.ID, response.Guess)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Game = updatedGame
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
