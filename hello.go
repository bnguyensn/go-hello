package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/google/go-cmp/cmp"

	"github.com/bnguyensn/go-hello/basics"
	"github.com/bnguyensn/go-hello/morestrings"
	"github.com/bnguyensn/go-hello/weirdstringcase"
)

func testBasics() {
basics.LolLoop(rand.Intn(15) + 5)
	fmt.Println("")
	basics.LmaoLoop(rand.Intn(15) + 5)

	basics.FizzBuzz(30)

	basics.CheckOs()

	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("More string"))
	fmt.Println(cmp.Diff("Hello world", "Hello Go"))

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10000)
	b := rand.Intn(10000)
	c := basics.Add(a, b)
	fmt.Println(a, " + ", b, " = ", c)

	testBasics();
}

func main() {
	s := weirdstringcase.WeirdString("This is a triumph")
	fmt.Println(s)
}
