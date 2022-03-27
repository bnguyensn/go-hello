package basics

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

// HelloWorld says 'Hello, world!'.
func HelloWorld() {
	fmt.Println("Hello, world!")
}

// HelloName says hello to a name. The hello message is selected randomly from a
// pre-defined list.
// If an empty name is provided, exit the program.
func HelloName(name string) string {
	if name == "" {
		log.Fatal("Please provide a non-empty name")
	}

	message := fmt.Sprintf(randomHelloFormat(), name)
	fmt.Println(message)

	return message
}

// HelloNames says hello to multiple names. The hello message is selected
// randomly from a pre-defined list.
// If an empty name list is provided, exit the program.
func HelloNames(names []string) map[string]string {
	if len(names) == 0 {
		log.Fatal("Please provide a non-empty name list")
	}

	messages := make(map[string]string)

	for _, name := range names {
		message := HelloName(name)

		messages[name] = message
	}

	return messages
}

// randomHelloFormat returns a randomly-selected hello message format.
func randomHelloFormat() string {
	rand.Seed(time.Now().UnixNano())

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v.",
		"Well met, %v.",
	}

	return formats[rand.Intn(len(formats))]
}

// Log some texts.
func Log(text string) {
	fmt.Println(text)
}

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
