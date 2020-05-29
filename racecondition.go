package main

import (
	"fmt"
	"sync"
)

//This application is an illustration for race condition
//Try to run the application several times and you probably have different result each time
//There are 2 processes one adds the number one decreases the number of a common integer
//Both of the process want to access the same resource which creates race condition
//By default we would expect that the cnt variable should be zero in the end, but it varies



//I added the wait group because to go program ends when the main function end,
//so it might not wait for the other goroutines to complete
var wg sync.WaitGroup

func main() {

	//shared resource cnt
	cnt := 0
	//the number of repetition, if it's higher the chance is higher to get different result than 0
	times := 50000
	wg.Add(1)
	go func() {
		for i:=0; i<times; i++ {
			cnt--   // try to comment this out and you get times every time
		}
		defer wg.Done()
	}()
	for i:=0; i<times; i++ {
		cnt++				// try to comment this out and you get -1*times every time
	}
	wg.Wait()
	fmt.Println(cnt)
}
