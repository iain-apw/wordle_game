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
