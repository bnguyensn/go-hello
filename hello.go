package main

import (
	"fmt"
	"time"

	"math/rand"
	"runtime"

	"github.com/bnguyensn/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func add(a, b int) int {
	return a + b
}

func lolLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("LOL")
	}
}

func lmaoLoop(count int) {
	i := 0
	for i < count {
		fmt.Printf("LMAO")
		i++
	}
}

func fizzBuzz(count int) {
	for i := 0; i < count; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz at ", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz at ", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz at ", i)
		}
	}
}

func checkOs() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("More string"))
	fmt.Println(cmp.Diff("Hello world", "Hello Go"))

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10000)
	b := rand.Intn(10000)
	c := add(a, b)
	fmt.Println(a, " + ", b, " = ", c)

	lolLoop(rand.Intn(15) + 5)
	fmt.Println("")
	lmaoLoop(rand.Intn(15) + 5)

	fizzBuzz(30)

	checkOs()
}
