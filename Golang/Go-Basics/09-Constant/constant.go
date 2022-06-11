package main

import (
	"fmt"
	"math"
)

func main() {

	r := 3.0
	const pi float64 = 3.14
	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)
	fmt.Println("")

	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)
	fmt.Println("")

	// const ----> compile time
	// var ----> run time

	var v = math.Pow(3, 4)
	fmt.Printf("%T, %v\n", v, v)

	// const c = math.Pow(3,4) // Initialized error

}
