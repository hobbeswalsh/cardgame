package cardgame

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func sum(a []int32) (s int32) {
	for _, v := range a {
		s += v
	}
	return
}

/*
	No, this isn't portable. Sorry!
*/

var Random *os.File

func init() {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	Random = f
}

func uuid() string {
	init()
	b := make([]byte, 16)
	Random.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func Get_uuid() string {
	return uuid()
}

func Get_winner_list(c chan []int32) {
	base_array := []int32{}
	base_array = append(base_array, get_list_of_nums(get_small_random_num, 15)...)
	base_array = append(base_array, get_list_of_nums(get_small_negative_random_num, 10)...)
	for sum(base_array) <= 0 {
		Get_winner_list(c)
	}
	c <- base_array
}

func Get_loser_list(c chan []int32) {
	base_array := []int32{}
	base_array = append(base_array, get_list_of_nums(get_big_random_num, 10)...)
	base_array = append(base_array, get_list_of_nums(get_big_negative_random_num, 15)...)
	for sum(base_array) >= 0 {
		Get_loser_list(c)
	}
	c <- base_array

}

func get_big_random_num() int32 {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return R.Int31n(1500) + 350
}

func get_big_negative_random_num() int32 {
	return -1 * get_big_random_num()
}

func get_small_random_num() int32 {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return R.Int31n(450) + 35
}

func get_small_negative_random_num() int32 {
	return -1 * get_small_random_num()
}

func get_list_of_nums(f func() int32, l int) []int32 {
	base_array := []int32{}
	for i := 0; i < l; i++ {
		base_array = append(base_array, f())
	}
	return base_array
}

func Shuffle_int32s(xs []int32) []int32 {
	for i := range xs {
		randchoice := R.Intn(len(xs))
		xs[i], xs[randchoice] = xs[randchoice], xs[i]
	}
	return xs
}
