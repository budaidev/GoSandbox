package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	var name string
	var address string
	m := make(map[string]string)

	fmt.Println("Please enter the name")
	fmt.Scan(&name)
	fmt.Println("Please enter the address")
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	res, _ := json.Marshal(m)
	fmt.Printf(string(res))
}