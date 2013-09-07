package cardgame

import (
	"errors"
	"fmt"
)

func getGame(s string) (CardGame, error) {
	return GetGameByUUID(s)
}

func MakePlay(s string, i int) (int32, error) {
	cg, err := getGame(s)
	if err != nil {
		return 0, err
	}
	if i > len(cg.Decks)-1 || i < 0 {
		return 0, errors.New("Invalid deck number")
	}
	deck := cg.Decks[i]
	card := deck.Values[0]
	deck.Values = deck.Values[1:]
	cg.Decks[i] = deck
	if len(deck.Values) == 0 {
		fmt.Println("out of cards!")
		cg.Regenerate()
		SaveGame(cg)
		return card, nil
	}

	SaveGame(cg)
	return card, nil
}

func NewGame() (CardGame, error) {
	cg, err := MakeNewGame()
	if err != nil {
		return CardGame{}, errors.New("No way Jose")
	}
	return cg, nil
}
