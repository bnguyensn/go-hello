package main

import (
	"github.com/bnguyensn/go-hello/greetings"
)

func main() {
	names := []string{"Alice", "Bob", "Claire", "David"}

	greetings.HelloNames(names)
}
