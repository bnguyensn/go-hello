package greetings

import (
	"fmt"
	"log"
	"math/rand"
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
