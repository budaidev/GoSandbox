package main

import (
	"fmt"
)

type Animals struct {
	food string
	locomotion string
	noise string
}

func (a *Animals) Eat(){
	fmt.Println(a.food)
}

func (a *Animals) Move(){
	fmt.Println(a.locomotion)
}

func (a *Animals) Speak(){
	fmt.Println(a.noise)
}

func main() {

	fmt.Println("Type the animal and the action separated by space!")

	m := make(map[string]Animals)
	var anim string
	var action string
	m["cow"] = Animals{food:"grass", locomotion:"walk",noise:"moo"}
	m["bird"] = Animals{food:"worms", locomotion:"fly",noise:"peep"}
	m["snake"] = Animals{food:"mice", locomotion:"slither",noise:"hsss"}

	for {
		fmt.Println(">")
		fmt.Scanln(&anim, &action)
		a, isin := m[anim]
		if isin {
			switch action {
				case "eat": a.Eat()
				case "move": a.Move()
				case "speak": a.Speak()
				default: fmt.Println("Unknown action")
			}
		} else {
			fmt.Println("Unknown animal")
		}
	}
}