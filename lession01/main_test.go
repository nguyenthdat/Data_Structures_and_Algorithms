package main

import "testing"

func TestLocateCard(t *testing.T) {
	cards := []int{13, 11, 12, 7, 4, 3, 1, 0}
	if locateCard(cards, 7) != 3 {
		t.Errorf("Something went wrong")
	}
}
