package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	if n < 1 {
		usage()
	}

	fmt.Print(1)
	for i := 2; i <= n; i++ {
		fmt.Print(" ", fizzbuzz(i))
	}
	fmt.Println()
	fmt.Print(1)
	for i := 2; i <= n; i++ {
		fmt.Print(" ", fizzbuzz2(i))
	}
}

func usage() {
	fmt.Println("usage:\n\tfizzbuzz <count>\n\n\tcount must be greater than 0.")
	os.Exit(1)
}

func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "fizzbuzz"
	case n%5 == 0:
		return "buzz"
	case n%3 == 0:
		return "fizz"
	default:
		return strconv.Itoa(n)
	}
}

// https://philcrissman.net/posts/eulers-fizzbuzz/
func fizzbuzz2(n int) string {
	m := [...]string{
		0:  "fizzbuzz",
		6:  "fizz",
		10: "buzz",
	}
	m[1] = strconv.Itoa(n)
	i := int(math.Mod(math.Pow(float64(n), 4), 15))
	return m[i]
}
