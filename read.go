package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}

func main() {
	persons := make([]Person, 0, 10)
	var filePath string
	fmt.Println("Please enter the file path")
	fmt.Scan(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		persons = append(persons, Person{s[0], s[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range persons{
		fmt.Println(p.firstName, p.lastName)
	}
}