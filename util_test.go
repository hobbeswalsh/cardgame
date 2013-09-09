package cardgame

import (
	"testing"
)

func TestSumZeros(t *testing.T) {
	arr := []int32{0, 0, 0, 0, 0}
	if sum(arr) != 0 {
		t.Error("Cannot sum 0-sum array")
	}
}

func TestSumNegatives(t *testing.T) {
	arr := []int32{-1, -2, -3}
	if sum(arr) != -6 {
		t.Error("Cannot sum negative array")
	}
}

func TestSumPositives(t *testing.T) {
	arr := []int32{1, 2, 3}
	if sum(arr) != 6 {
		t.Error("Cannot sum positive array")
	}
}

func TestGetUuid(t *testing.T) {
	uuid := Get_uuid()
	if len(uuid) != 36 {
		t.Error("Could not get UUID from system call")
	}
}

func TestGetWinnerListReturnsWinner(t *testing.T) {
	c := make(chan []int32)
	go Get_winner_list(c)
	card_list := <-c
	if sum(card_list) < 0 {
		t.Error("Winner list isn't really a winner...")
	}
}
