package main

import (
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = [...]struct {
		n    int
		want string
	}{
		{1, "1"},
		{5, "buzz"},
		{-5, "buzz"},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			got := fizzbuzz(test.n)
			if got != test.want {
				t.Fatal("got:", got, "want:", test.want)
			}
		})
	}
}
