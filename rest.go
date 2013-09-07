package cardgame

import (
	"encoding/json"
	"fmt"
	"github.com/bpowers/seshcookie"
	"index/suffixarray"
	"net/http"
	"strconv"
)

func throw_json_error(rw http.ResponseWriter, errorstr string) {
	rw.WriteHeader(400)
	rw.Header().Set("Content-Type", "text/json")
	err := make(map[string]interface{})
	err["error"] = errorstr
	err["success"] = false
	j, _ := json.Marshal(err)
	rw.Write(j)
}

type RootHandler struct{}

func (h *RootHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if check_url("/", req.URL.Path) == false {
		return
	}

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

type PlayHandler struct{}

func (h *PlayHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !check_url("/play", req.URL.Path) {
		return
	}
	session := seshcookie.Session.Get(req)

	uuid, _ := session["uuid"].(string)

	deck := req.FormValue("deck")
	if deck == "" {
		throw_json_error(rw, uuid)
		return
	}

	throw_json_error(rw, "can't use this endpoint yet.")
	return
	// // func (r *Request) PostFormValue(key string) string

	// rw.WriteHeader(200)
	// rw.Header().Set("Content-Type", "text/plain")
	// var output []byte
	// strconv.AppendQuoteToASCII(output, "Well hello there!\n")

	// rw.Write(output)
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
