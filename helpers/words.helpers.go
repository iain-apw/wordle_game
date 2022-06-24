package helpers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func GenerateWord(letters int) (string, error) {
	fileName := fmt.Sprintf("./data/%v_words.txt", letters)

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var words []string

	s := bufio.NewScanner(file)
	for s.Scan() {
		words = append(words, s.Text())
	}

	randomIndex := rand.Intn(len(words))
	chosenWord := words[randomIndex]

	return chosenWord, nil
}
