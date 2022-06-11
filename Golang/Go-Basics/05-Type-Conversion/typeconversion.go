package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)
	// fmt.Println(x + y) -> mismatched ind and float64 error.

	// Type Conversion -> type(value) => int(y)
	fmt.Println(x + int(y))
	fmt.Println(float64(x) + y)

	fmt.Println("*********************************************")

	var a int8 = 10
	var b int16 = 10
	fmt.Println(a + int8(b))

	var i int8 = 127
	var j int16
	j = int16(i)
	fmt.Println(j)

	var k int16 = 128
	var f int8
	f = int8(k)
	println(f)

	fmt.Println("*********************************************")

	// String + Number

	n := 10
	s := "10"

	fmt.Printf("%v %T \n", n, n)
	fmt.Printf("%v %T \n", s, s)
	fmt.Println("*********************************************")

	// fmt.Println(n + s) // => cannot convert type string to type int error

	num1 := 106
	str1 := string(rune(num1))

	fmt.Printf("%v %T \n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T \n", str1, str1)

}
