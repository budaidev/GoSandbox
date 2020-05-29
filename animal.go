package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (a Cow) Eat(){
	fmt.Println("grass")
}
func (a Cow) Move(){
	fmt.Println("walk")
}
func (a Cow) Speak(){
	fmt.Println("moo")
}

type Bird struct {
}

func (a Bird) Eat(){
	fmt.Println("worms")
}
func (a Bird) Move(){
	fmt.Println("fly")
}
func (a Bird) Speak(){
	fmt.Println("peep")
}

type Snake struct {
}

func (a Snake) Eat(){
	fmt.Println("mice")
}
func (a Snake) Move(){
	fmt.Println("slither")
}
func (a Snake) Speak(){
	fmt.Println("hsss")
}

func main() {

	fmt.Println("Create new animal with the newanimal command (e.g.):")
	fmt.Println("newanimal milka cow")
	fmt.Println("Query existing animals with the comman query (e.g):")
	fmt.Println("query milka eat")

	m := make(map[string]Animal)
	var command string
	var param1 string
	var param2 string
	for {
		fmt.Println(">")
		fmt.Scanln(&command, &param1, &param2)
		if command == "newanimal" {
			error := false
			switch param2 {
				case "cow": m[param1] = Cow{}
				case "snake": m[param1] = Snake{}
				case "bird": m[param1] = Bird{}
				default: {fmt.Println("Animal type does not exist (only cow, snake or bird allowed")
					error = true}
			}
			if !error{
				fmt.Println("Created it!")
			}
		} else if command == "query" {
			a, isin := m[param1]
			if isin {
				switch param2 {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Unknown action")
				}
			} else {
				fmt.Println("Animal not found")
			}
		}
	}
}