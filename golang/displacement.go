package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	return func(t float64) float64 {
		var finalDisplacement float64
		finalDisplacement = (1/2*acceleration)*(t*t) + ((velocity * t) + displacement)
		return finalDisplacement
	}

}

func main() {
	var acceleration, velocity, displacement float64
	var t float64

	for {
		fmt.Printf("Enter acceleration: ")
		fmt.Scan(&acceleration)
		fmt.Print("Enter velocity: ")
		fmt.Scan(&velocity)
		fmt.Print("Enter displacement: ")
		fmt.Scan(&displacement)
		fn := GenDisplaceFn(acceleration, velocity, displacement)
		fmt.Printf("Enter Time ")
		fmt.Scan(&t)
		fmt.Println(fn(t))
	}

}
