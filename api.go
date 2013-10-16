package cardgame

import (
	"errors"
)

func getGame(s string) (CardGame, error) {
	return GetGameByUUID(s)
}

func MakePlay(s string, i int) (int32, CardGame, error) {
	cg, err := getGame(s)
	if err != nil {
		return 0, CardGame{}, err
	}
	if i > len(cg.Decks)-1 || i < 0 {
		return 0, CardGame{}, errors.New("Invalid deck number")
	}
	deck := cg.Decks[i]
	card := deck.Values[0]
	deck.Values = deck.Values[1:]
	cg.Decks[i] = deck
	if deck.Risky() {
		// do nothing now, but at a later date, log a risky play.
	} else {
		// do nothing now, but at a later date, log a safe play.
	}
	if len(deck.Values) == 0 {
		cg.Regenerate()
		SaveGame(cg)
		return card, CardGame{}, nil
	}

	SaveGame(cg)
	cg, _ = getGame(s)
	return card, cg, nil
}

func NewGame() (CardGame, error) {
	cg, err := MakeNewGame()
	if err != nil {
		return CardGame{}, errors.New("No way Jose")
	}
	return cg, nil
}
