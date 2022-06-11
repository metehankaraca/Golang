// Addition + => 15 + 10 operant, + operator
// Substruction -
// Product *
// Division /
// Remainder %

package main

import "fmt"

func main() {
	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Println("")
	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x - y), (x - y))
	fmt.Printf("%T, %v\n", (x * y), (x * y))
	fmt.Printf("%T, %v\n", (x / y), (x / y))
	fmt.Printf("%T, %v\n", (x % y), (x % y))
	fmt.Println("")

	z := float64(5 / 2)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Println("")

	f := 5.0 / 2
	fmt.Printf("%T, %v\n", f, f)
	fmt.Println("")

	// Increment ++, Decrement -- // Postfix

	j := 10
	fmt.Println(j)
	j = j + 1
	fmt.Println(j)
	// fmt.Println(j++) // unexpected ++ error
	j++
	fmt.Println(j)
	fmt.Println("")

	k := 10
	fmt.Println(k)
	k = k - 1
	fmt.Println(k)
	// fmt.Println(k--) // unexpected ++ error
	k--
	fmt.Println(k)

}
