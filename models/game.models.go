package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iain-apw/wordle_game/helpers"
)

type GameStatus string

const (
	GameStatus_Won        GameStatus = "Won"
	GameStatus_Lost       GameStatus = "Lost"
	GameStatus_InProgress GameStatus = "InProgress"
)

type Game struct {
	ID         string     `json:"id"`
	Letters    int        `json:"letters"`
	UserId     string     `json:"userId"`
	Answer     string     `json:"answer"`
	Status     GameStatus `json:"status,omitempty"`
	Guesses    []Guess    `json:"guesses,omitempty"`
	CreatedAt  time.Time  `json:"createdAt,omitempty"`
	LastPlayed time.Time  `json:"lastPlayed,omitempty"`
}

func NewGame(letters int, user *User) Game {

	// Generate a new word
	word, err := helpers.GenerateWord(letters)

	if err != nil {
		log.Fatal(err)
	}

	return Game{
		ID:        uuid.New().String(),
		UserId:    user.ID,
		Answer:    word,
		Letters:   letters,
		Status:    GameStatus_InProgress,
		CreatedAt: time.Now(),
	}
}

type Guess struct {
	Word string
	// Something to store the state of the guess
}

type CreateGameRequest struct {
	Letters int `json:"letters"`
}
