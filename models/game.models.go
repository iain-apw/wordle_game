package models

import (
	"time"
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

type Guess struct {
	Word string `json:"word"`
	// Something to store the state of the guess
	Letters []GuessLetter `json:"letters"`
	Correct bool          `json:"correct"`
}

type LetterStatus string

const (
	LetterStatus_Missing    LetterStatus = "Missing"
	LetterStatus_WrongPlace LetterStatus = "WrongPlace"
	LetterStatus_Correct    LetterStatus = "Correct"
)

type GuessLetter struct {
	Letter string       `json:"letter"`
	Status LetterStatus `json:"status"`
}

type CreateGameRequest struct {
	Letters int `json:"letters"`
}

type MakeGuessRequest struct {
	Guess string `json:"guess"`
}

type MakeGuessResponse struct {
	IsValidWord bool  `json:"isValidWord"`
	Guess       Guess `json:"guess,omitempty"`
	Game        Game  `json:"game,omitempty"`
}
