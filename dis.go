package main

import (
	"fmt"
)

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter Acceleration:")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Initial Velocity:")
	fmt.Scanf("%f\n", &v0)
	fmt.Println("Initial Displacement:")
	fmt.Scanf("%f\n", &s0)
	dfunc := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter Time: ")
	fmt.Scanf("%f\n", &t)
	fd := dfunc(t)
	fmt.Printf("Displacement in time %.2f seconds : %.2f\n", t, fd)
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return ((a*(t*t))/2 + (v0 * t) + s0)
	}
}
