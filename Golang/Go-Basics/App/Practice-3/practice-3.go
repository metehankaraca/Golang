package main

import "fmt"

func main() {
	// 1-) x = x - 10 vs x -= 10

	x := 50
	// x = x - 10 // Assignment statement
	x -= 10 // Assignment operation

	fmt.Printf("%T, %v\n", x, x)
	fmt.Println("")

	//
	main2()
	main3()
	main4()
	main5()
	main6()
}

func main2() {
	// 2-) K = F - 32 / 1.8 + 273 => -40 Fahreneit to kelvin

	F := -40
	K := float64(F-32)/1.8 + 273

	fmt.Printf("%T, %v\n", K, K)
	fmt.Println("")
}

func main3() {
	const myAge = 40

	fmt.Printf("%v, %T\n", myAge, myAge)
	fmt.Println("")
}

// main4
const s = 24

func main4() {
	// Constant with shadowing
	const s = 14
	fmt.Printf("%T, %v\n", s, s)
	fmt.Println("")

}

func main5() {
	// const x = 4, const y = 5.4 => x + y ?
	const x = 4   //typeless
	const y = 5.4 // typeless

	fmt.Printf("%T, %v\n", x+y, x+y) // x data type float64
	fmt.Println("")
}

func main6() {
	// const x float64 = 6.4, y := 4 + x, y ?
	const x float64 = 6.4
	y := 4 + x

	fmt.Printf("%T, %v\n", y, y)
}
