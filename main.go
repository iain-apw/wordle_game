package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iain-apw/wordle_game/handlers"
)

func main() {
	fmt.Println("Hello there Go Wordle Game!")

	router, err := handlers.InitApiRoutes()

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}

	log.Fatal(http.ListenAndServe(":8000", router))
}

/*
	TODO: REMOVE THIS ONCE WE HAVE THE PROCESSING IN PLACE
	fmt.Println("WORD PROCESSING!")

	const fileName = "./data/all_words.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error occurred opening file %w", err)
		return
	}

	defer file.Close()

	//var words4, words5, words6, words7, words8, words9 []string

	wordsByLength := make(map[int][]string)

	count := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		word := s.Text()
		wordsByLength[len(word)] = append(wordsByLength[len(word)], word)
		count++
	}

	fmt.Printf("File contains %v words\n", count)

	fmt.Printf("4 letter words: %v\n", len(wordsByLength[4]))
	fmt.Printf("5 letter words: %v\n", len(wordsByLength[5]))
	fmt.Printf("6 letter words: %v\n", len(wordsByLength[6]))
	fmt.Printf("7 letter words: %v\n", len(wordsByLength[7]))
	fmt.Printf("8 letter words: %v\n", len(wordsByLength[8]))
	fmt.Printf("9 letter words: %v\n", len(wordsByLength[9]))

	for key, element := range wordsByLength {
		fmt.Println("Key:", key, "=>", "Element:", len(element))

		newFileName := fmt.Sprintf("./data/%v_words.txt", key)

		wordData := strings.Join(element, "\n")

		fmt.Println(newFileName)

		f, err := os.Create(newFileName)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		n2, err := f.WriteString(wordData)
		if err != nil {
			panic(err)
		}

		fmt.Printf("wrote %d bytes\n", n2)
	}

	return
*/
