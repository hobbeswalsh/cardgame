package cardgame

import (
	"encoding/json"
	"fmt"
	"github.com/bpowers/seshcookie"
	"index/suffixarray"
	"net/http"
	"strconv"
)

type PlayResult struct {
	ChosenCard    int32    `json:chosenCard`
	GameRemaining CardGame `json:gameRemaining`
}

func throw_json_error(rw http.ResponseWriter, errorstr string) {
	rw.WriteHeader(400)
	rw.Header().Set("Content-Type", "text/json")
	err := make(map[string]interface{})
	err["error"] = errorstr
	err["success"] = false
	j, _ := json.Marshal(err)
	rw.Write(j)
}

func return_json_success(rw http.ResponseWriter, i interface{}) {
	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "text/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	j, _ := json.Marshal(i)
	rw.Write(j)
}

/*
  RootHandler handles requests to fetch whole games.
*/
type RootHandler struct{}

func (h *RootHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	session := seshcookie.Session.Get(req)

	uuid, _ := session["uuid"].(string)
	if uuid == "" {
		cg, err := MakeNewGame()
		if err == nil {
			session["uuid"] = cg.Uuid
			rw.WriteHeader(200)
			rw.Header().Set("Content-Type", "text/json")
			j, err := json.Marshal(cg)
			if err != nil {
				throw_json_error(rw, "Couldn't marshal new cardgame into JSON")
				return
			} else {
				rw.Write(j)
			}
		}
		return
	}
	cg, err := GetGameByUUID(uuid)
	if err != nil {
		errtext := fmt.Sprintf("Could not fetch game by UUID %s", uuid)
		throw_json_error(rw, errtext)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "text/json")

	j, err := json.Marshal(cg)
	if err != nil {
		errtext := fmt.Sprintf("Couldn't marshal cardgame %s into JSON", uuid)
		throw_json_error(rw, errtext)
		return
	} else {
		rw.Write(j)
	}

}

/*
 PlayHandler handles plays.
*/
type PlayHandler struct{}

func (h *PlayHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	session := seshcookie.Session.Get(req)

	uuid, _ := session["uuid"].(string)

	deck := req.FormValue("deck")
	if deck == "" {
		throw_json_error(rw, uuid)
		return
	}

	deck_num, err := strconv.ParseInt(deck, 10, 0)
	if err != nil {
		throw_json_error(rw, fmt.Sprintf("Couldn't convert %s to an integer", deck))
		return
	}
	card_chosen, remaining_game, err := MakePlay(uuid, int(deck_num))
	if err != nil {
		throw_json_error(rw, fmt.Sprintf("Could not make a play for some reason: %s", err))
		return
	}
	play := PlayResult{card_chosen, remaining_game}
	return_json_success(rw, play)
}

func check_url(substr, str string) bool {
	var dst []byte
	substr_bytes := strconv.AppendQuoteToASCII(dst, substr)
	str_bytes := strconv.AppendQuoteToASCII(dst, str)
	index := suffixarray.New(str_bytes)
	offsets := index.Lookup(substr_bytes, -1)
	if offsets == nil {
		return false
	}
	return offsets[0] == 0
}
