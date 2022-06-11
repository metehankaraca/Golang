// Comparison operators
// == Equal != Not equal
// < Less than > Greater than
// <= Less than or equal >= Greater than or equal

// Logical Operators
// && conditional AND  p && q  is "if p then q else false"
// || conditional OR  p || q is "if p then true else q"
// ! NOT   !p is "not p"

package main

import "fmt"

func main() {
	// Comparison Operators
	// comparison can only be made on the same data types !!

	x, y := 3, 7
	// x, y := 3, 7.0  => invalid operation: mismatched types ind and float64 error.
	// x, y := 3, true  => invalid operation: mismatched types ind and float64 error.
	// x, y := 3, "b"  => invalid operation: mismatched types ind and float64 error.

	fmt.Printf("%T, %v\n", x == y, x == y) // false
	fmt.Printf("%T, %v\n", x != y, x != y) // true
	fmt.Printf("%T, %v\n", x < y, x < y)   // true
	fmt.Printf("%T, %v\n", x > y, x > y)   // false
	fmt.Printf("%T, %v\n", x >= y, x >= y) // false
	fmt.Printf("%T, %v\n", x <= y, x <= y) // true
	fmt.Println("")

	// Logical Operators
	q, w := 15, 20

	set1 := (q == w) // false
	set2 := (q < w)  // true

	fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)
	fmt.Println("")

	fmt.Printf("%v\n", (set1 && set2)) // false && true --> false // returns true if both conditions are met
	fmt.Printf("%v\n", (set1 || set2)) // false && true --> true  // the result will be true if one of the two conditions returns true

}
