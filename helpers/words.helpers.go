package helpers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/iain-apw/wordle_game/models"
)

func readAllWords(letters int) ([]string, error) {
	fileName := fmt.Sprintf("./data/%v_words.txt", letters)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var words []string

	s := bufio.NewScanner(file)
	for s.Scan() {
		words = append(words, strings.ToUpper(s.Text()))
	}

	return words, nil
}

func GenerateWord(letters int) (string, error) {
	words, err := readAllWords(letters)

	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(words))
	chosenWord := words[randomIndex]

	return chosenWord, nil
}

func IsValidWord(word string, letters int) (bool, error) {
	length := len(word)

	if length != letters {
		return false, nil
	}

	words, err := readAllWords(length)

	if err != nil {
		return false, err
	}

	comparisonWord := strings.ToUpper(strings.TrimSpace(word))
	matched := false

	for _, w := range words {
		if w == comparisonWord {
			// Found!
			matched = true
			break
		}
	}

	return matched, err
}

func CheckGuess(guess string, answer string) models.Guess {

	if len(guess) == len(answer) {

		letters := make([]models.GuessLetter, len(guess))

		upperGuess := strings.ToUpper(guess)

		for i := 0; i < len(guess); i++ {
			if upperGuess[i] == answer[i] {
				letters[i] = models.GuessLetter{
					Letter: string(upperGuess[i]),
					Status: models.LetterStatus_Correct,
				}
			}
		}

		for i := 0; i < len(guess); i++ {
			if letters[i].Letter == "" {
				found := false

				for j := 0; j < len(guess); j++ {
					if i != j && letters[j].Letter == "" && upperGuess[i] == answer[j] {
						found = true
						break
					}
				}

				if found {
					letters[i] = models.GuessLetter{
						Letter: string(upperGuess[i]),
						Status: models.LetterStatus_WrongPlace,
					}
				}
			}
		}

		for i := 0; i < len(guess); i++ {
			if letters[i].Letter == "" {
				letters[i] = models.GuessLetter{
					Letter: string(upperGuess[i]),
					Status: models.LetterStatus_Missing,
				}
			}
		}

		return models.Guess{
			Word:    upperGuess,
			Letters: letters,
			Correct: upperGuess == answer,
		}

	}

	return models.Guess{}
}
