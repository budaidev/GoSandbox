package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func (float64) float64 {
	return func(t float64) float64{
		return a/2*t*t+v0*t+s0
	}
}


func main() {
	fmt.Println("Please enter acceleration")
	var a float64
	fmt.Scan(&a)
	fmt.Println("Please enter initial velocity")
	var v0 float64
	fmt.Scan(&v0)
	fmt.Println("Please enter initial displacement")
	var s0 float64
	fmt.Scan(&s0)
	var in string



	fn := GenDisplaceFn(a, v0, s0)
	for{
		fmt.Println("Please enter time (or X to exit)")
		fmt.Scan(&in)
		if in == "X" || in == "x"{
			break
		}
		time, _ := strconv.ParseFloat(in, 64)
		fmt.Println(fn(time))
	}

}