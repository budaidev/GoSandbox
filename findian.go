package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	i := unicode.ToLower(rune(input[0]))
	n := unicode.ToLower(rune(input[len(input)-1]))
	if i =='i' &&
		n=='n' &&
		strings.Index(input, "a")>0{
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}