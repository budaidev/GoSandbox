package main

import (
	"fmt"
)

// GenDisplaceFn returns a time dependent function.
func GenDisplaceFn2(acc float64, velInit float64, disInit float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return 0.5*acc*t*t + velInit*t + disInit
	}

	return fn
}

func main() {

	var acceleration float64 = 0.0
	var initialVel float64 = 0.0
	var initialDisp float64 = 0.0
	var duration float64 = 0.0

	fmt.Printf("Enter a value for acceleration: ")
	_, err1 := fmt.Scan(&acceleration)

	fmt.Printf("Enter a value for the initial velocity: ")
	_, err2 := fmt.Scan(&initialVel)

	fmt.Printf("Enter a value for the initial displacement: ")
	_, err3 := fmt.Scan(&initialDisp)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("An unknown error occurred.")
	}

	fn := GenDisplaceFn2(acceleration, initialVel, initialDisp)

	fmt.Printf("Now enter a time duration: ")
	_, err4 := fmt.Scan(&duration)

	if err4 != nil {
		fmt.Println("An unknown error occurred.")
	}

	fmt.Printf("The total displacement is %f", fn(duration))
}
