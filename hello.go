package main

import (
	"fmt"
)

type BasicInfo struct {
	name string
	sex  string
}

type Character struct {
	self      BasicInfo
	father    BasicInfo
	mother    BasicInfo
	isAlive   bool
	age       int
	responses []string
}

func main() {
	charPaul := Character{
		self: BasicInfo{
			name: "Paul Atreides",
			sex:  "male",
		},
		father: BasicInfo{
			name: "Leto Atreides",
			sex:  "male",
		},
		mother: BasicInfo{
			name: "Jessica",
			sex:  "female",
		},
		isAlive:   true,
		age:       16,
		responses: []string{"Hi there!", "Greetings!", "Nice to meet you!"},
	}
	charPaulPointer := &charPaul

	fmt.Printf("The first character's name is %v.\n"+
		"Their sex is %v.\n"+
		"Their father is %v.\n"+
		"Their mother is %v.\n"+
		"*****\n"+
		"ALIVE: %v.\n"+
		"AGE: %v.\n",
		charPaul.self.name,
		charPaul.self.sex,
		charPaul.father.name,
		charPaul.mother.name,
		charPaul.isAlive,
		charPaul.age,
	)

	fmt.Println("*****")

	for _, response := range charPaul.responses {
		fmt.Printf("%v says: %v\n",
			charPaul.self.name,
			response,
		)
	}

	fmt.Println("*****")

	fmt.Printf("charPaulPointer: %v\n", &charPaulPointer)
}
