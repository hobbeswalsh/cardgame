package cardgame

import (
	"log"
	"math/rand"
	"os/exec"
	"time"

)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func sum(a []int32) (s int32) {
	for _, v := range a {
		s += v
	}
	return
}

func Get_uuid() string {
	out, err := exec.Command("/usr/bin/uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:len(out)-1])
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
