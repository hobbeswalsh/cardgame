package main

import (
	"github.com/hobbeswalsh/cardgame"
	"log"
	"net/http"
)

func main() {
	cardgame.CreateRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
