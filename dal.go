package cardgame

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var url = "mongodb://localhost"
var session, err = mgo.Dial(url)

func get_session() *mgo.Collection {
	c := session.DB("cardgame").C("games")
	return c
}

func GetGameByUUID(s string) (CardGame, error) {
	c := get_session()
	var result CardGame
	err := c.Find(bson.M{"uuid": s}).One(&result)
	return result, err

}

func MakeNewGame() (CardGame, error) {
	c := get_session()
	cg := MakeCardgame()
	err := c.Insert(cg)
	if err != nil {
		fmt.Println(err)
	}
	return cg, err
}

func SaveGame(cg CardGame) bool {
	c := get_session()
	change := mgo.Change{
		Update:    bson.M{"decks": cg.Decks, "uuid": cg.Uuid},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": cg.Id}).Apply(change, &cg)
	if err != nil {
		fmt.Println("Could not update cardgame!")
		return false
	}
	return true
}
