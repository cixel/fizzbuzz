package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("🤬")
	}

	for i := 1; i <= n; i++ {
		fmt.Println(fizzbuzz(i))
	}
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

// First line is "mode: foo", where foo is "set", "count", or "atomic".
// Rest of file is in the format
//	encoding/base64/base64.go:34.44,37.40 3 1
// where the fields are: name.go:line.column,line.column numberOfStatements count
//
// mode: set
// ehden.net/fizzbuzz/fizzbuzz.go:10.13,13.16 3 0
// ehden.net/fizzbuzz/fizzbuzz.go:13.16,15.3 1 0
// ehden.net/fizzbuzz/fizzbuzz.go:17.2,17.26 1 0
// ehden.net/fizzbuzz/fizzbuzz.go:17.26,19.3 1 0
// ehden.net/fizzbuzz/fizzbuzz.go:22.29,23.9 1 1
// ehden.net/fizzbuzz/fizzbuzz.go:24.17,25.20 1 0
// ehden.net/fizzbuzz/fizzbuzz.go:26.16,27.16 1 1
// ehden.net/fizzbuzz/fizzbuzz.go:28.16,29.16 1 0
// ehden.net/fizzbuzz/fizzbuzz.go:30.10,31.25 1 1
// ehden.net/fizzbuzz/random_other_file.go:5.21,7.2 1 0

// 10:13, 12:16 --> 10:14, 12:17
// 16:2, 16:26 --> 16:3, 16:27
// 12:16, 14:3 --> 12:17, 14:4
// 16:26, 18:3 --> 16:27, 18:4
// 21:29, 22:9 --> 21:30, 22:10
// 23:17, 24:20 --> 23:18, 24:21
// 25:16, 26:16 --> 25:17, 26:17
// 27:16, 28:16 --> 27:17, 28:17
// 29:10, 30:25 --> 29:11, 30:26
