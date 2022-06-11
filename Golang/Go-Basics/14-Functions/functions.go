package main

import (
	"fmt"
)

func main() {
	// var x, y, sum int
	// x = 5
	// y = 10
	// sum = x + y
	// fmt.Printf("%d ve %d toplami %d\n", x, y, sum)

	// Moduler Programming
	// Readable Code
	// From complex to simple

	sum(5, 10)
	fmt.Println(sum(5, 15))

	hello()

	// return vs print
	sum(5, 10)   // not written to console
	sum2(10, 10) // written to console
	fmt.Println("")

	fmt.Println(result(60))
	fmt.Println(result(40))
	fmt.Println("")

}

// func <funcName>(parameters) returntype { code }

func hello() {
	fmt.Println("Hello Basic Func")
}

func sum(x int, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func result(grade int) string {
	if grade >= 50 {
		return "Congrats!"
	}
	return "Sorry!"

}
