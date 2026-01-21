package main

import (
	"fmt"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (a Animal) Eat() {
	fmt.Println(a.eat)

}

func (a Animal) Move() {
	fmt.Println(a.move)
}

func (a Animal) Speak() {
	fmt.Println(a.speak)
}

var list = make(map[string]*Animal)

func predefined() {
	list["cow"] = &Animal{
		"grass",
		"walk",
		"moo",
	}
	list["bird"] = &Animal{
		"worms",
		"fly",
		"peep",
	}
	list["snake"] = &Animal{
		"mice",
		"slither",
		"hsss",
	}

}

func getUserInput() {
	var animalName string
	var information string
	predefined()
	for {
		fmt.Print("> ")
		fmt.Scan(&animalName, &information)
		animal := list[animalName]
		switch information {
		case "eat":
			animal.Eat()
		case "speak":
			animal.Speak()
		case "move":
			animal.Move()
		}

	}
}

func main() {
	getUserInput()
}
