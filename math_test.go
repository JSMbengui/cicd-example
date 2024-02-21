package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Error("Sum result is invalid. Received %d. Expected %d", total, 30)
	}
}