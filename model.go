package cardgame

import (
	"labix.org/v2/mgo/bson"
)

type Deck struct {
	Id     bson.ObjectId `bson:"_id" json:"-"`
	Values []int32       `bson:"values" json:"values"`
}

type CardGame struct {
	Id    bson.ObjectId `bson:"_id" json:"-"`
	Uuid  string        `bson:"uuid" json:"uuid"`
	Decks []Deck        `bson:"decks" json:"decks"`
}

func get_decks() []Deck {
	c := make(chan []int32)
	go Get_winner_list(c)
	go Get_winner_list(c)
	go Get_loser_list(c)
	go Get_loser_list(c)
	w, x, y, z := <-c, <-c, <-c, <-c
	d1 := Deck{bson.NewObjectId(), w}
	d2 := Deck{bson.NewObjectId(), x}
	d3 := Deck{bson.NewObjectId(), y}
	d4 := Deck{bson.NewObjectId(), z}
	return []Deck{d1, d2, d3, d4}
}

func MakeCardgame() CardGame {
	cg := CardGame{bson.NewObjectId(), Get_uuid(), get_decks()}
	cg.Shuffle()
	return cg
}

func (cg *CardGame) Regenerate() bool {
	cg.Decks = get_decks()
	cg.Shuffle()
	return true
}

func (cg *CardGame) Shuffle() bool {
	decks := cg.Decks
	for i := range decks {
		randchoice := R.Intn(len(decks))
		decks[i], decks[randchoice] = decks[randchoice], decks[i]
	}
	return true
}

func Shuffle_decks(xs []Deck) []Deck {
	for i := range xs {
		randchoice := R.Intn(len(xs))
		xs[i], xs[randchoice] = xs[randchoice], xs[i]
	}
	return xs
}
