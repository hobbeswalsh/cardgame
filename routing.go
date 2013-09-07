package cardgame

import (
	"github.com/bpowers/seshcookie"
	"net/http"
	"os"
)

func CreateRoutes() bool {
	key := make([]byte, 100)
	f, _ := os.Open("/dev/urandom")
	_, _ = f.Read(key)
	keyString := string(key)
	http.Handle("/", seshcookie.NewSessionHandler(
		&RootHandler{},
		keyString,
		nil))
	http.Handle("/play", seshcookie.NewSessionHandler(
		&PlayHandler{},
		keyString,
		nil))

	return true

}
