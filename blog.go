package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Blog struct {
	name, description string
}

var DB = new(BlogDB)
var BG = new(Blog)
var BC = new(BlogController)

func (b *Blog) Run() {
	loadENV()
	http.HandleFunc("/favicon.ico", BC.handleFavicon)
	http.HandleFunc("/", BC.handleRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (b *Blog) handleErrors(errtxt string) {
	panic(errtxt)
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		BG.handleErrors("Error loading .env file")
	}
}

func main() {
	BG.Run()
}
