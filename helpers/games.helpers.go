package helpers

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iain-apw/wordle_game/models"
)

func NewGame(letters int, user *models.User) models.Game {

	// Generate a new word
	word, err := GenerateWord(letters)

	if err != nil {
		log.Fatal(err)
	}

	return models.Game{
		ID:        uuid.NewString(),
		UserId:    user.ID,
		Answer:    word,
		Letters:   letters,
		Status:    models.GameStatus_InProgress,
		CreatedAt: time.Now(),
	}
}
