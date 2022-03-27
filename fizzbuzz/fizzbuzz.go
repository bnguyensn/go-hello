package fizzbuzz

import (
	"strconv"

	"github.com/bnguyensn/go-hello/basics"
)

// FizzBuzz prints 'Buzz' at indices divisible by 5, 'Fizz' at indices divisible
// by 3, and 'FizzBuzz' at indices divisible by both 5 and 3.
func FizzBuzz(count int) {
	for i := 0; i < count; i++ {
		if i%15 == 0 {
			basics.Log("FizzBuzz at " + strconv.Itoa(i))
		} else if i%5 == 0 {
			basics.Log("Buzz at " + strconv.Itoa(i))
		} else if i%3 == 0 {
			basics.Log("Fizz at " + strconv.Itoa(i))
		}
	}
}
