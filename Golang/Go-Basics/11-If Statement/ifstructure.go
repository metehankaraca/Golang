package main

import "fmt"

func main() {
	// If statement
	// if <boolean expression> { code }

	if !false {
		fmt.Println("Message 1")
	}

	if 5 > 3 {
		fmt.Println("Message 2")
	}

	fmt.Println("")

	x := 27

	// if <boolean expression> { code } else { code }

	if x%2 != 0 {
		fmt.Println(x, "even number")
	} else {
		fmt.Println(x, "odd number")
	}

	fmt.Println("")

	// if <boolean expression> { code } else if <boolean expression> { code } else { code }

	y := -5
	if y < 0 {
		fmt.Println(y, "negative number")
	} else if y%2 == 0 {
		fmt.Println(y, "even number")
	} else {
		fmt.Println(y, "odd number")
	}

	// Declaring variables in an if structure
	if j := 5; j > 2 {
		fmt.Println(j, "> 2")
	} else {
		fmt.Println(j, "< 2")
	}

}
