package basics

import (
	"fmt"
	"runtime"
)

// Adds 2 integers
func Add(a, b int) int {
	return a + b
}

// LolLoop prints count number of times the word "LOL".
func LolLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("LOL")
	}
}

// LmaoLoop prints count number of times the word "LMAO".
func LmaoLoop(count int) {
	i := 0
	for i < count {
		fmt.Printf("LMAO")
		i++
	}
}

// FizzBuzz prints 'Buzz' at indices divisible by 5, 'Fizz' at indices divisible by 3, and 'FizzBuzz' at indices divisible by both 5 and 3.
func FizzBuzz(count int) {
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

// CheckOs checks the OS that the go program runs on.
func CheckOs() {
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
