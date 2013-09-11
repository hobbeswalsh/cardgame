package cardgame

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

var (
	mgoSession   *mgo.Session
	databaseName = "myDB"
)

func get_session() mgo.Collection {
	var username = os.Getenv("CARDGAME_USER")
	var password = os.Getenv("CARDGAME_PASS")

	if mgoSession == nil {
		var url = fmt.Sprintf("mongodb://%s:%s@ds027738.mongolab.com:27738/cardgame", username, password)

		// var url = "mongodb://cgclient:c4rdg4m3@ds027738.mongolab.com:27738/cardgame"
		var err error
		fmt.Println(url)
		mgoSession, err = mgo.Dial(url)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	sesh := mgoSession.Clone()
	return *sesh.DB("cardgame").C("games")

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
