// 1 : int x, float64 y type conversion sample

package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 75
	var y float64
	y = float64(x)

	fmt.Println(y)

	//
	main1()
	main2()
	main3()
	main4()
}

// 2 : multiple assing sample x, y = y, x

func main1() {
	x := 5
	y := 10

	fmt.Println("X:", x, "Y", y)
	x, y = y, x
	fmt.Println("X:", x, "Y", y)

}

// 3 : non English variable names

func main2() {
	yaş := 40
	fmt.Println(yaş) // Turkish

	名稱 := "Name"
	яш := 30
	fmt.Println(名稱, яш) // Chinese - Russian

	// but since the traditional software language is English, the variable names should be named in English
}

// 4 : shadowing concept ?

func main3() {
	x := 5

	if true {
		x = 10
		x++
		fmt.Println(x)
	}

	fmt.Println(x)
}

// 5 : 65 as a string

func main4() {
	x := 65
	s := string(rune(x))

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", s, s)

	y := strconv.Itoa(x)
	fmt.Printf("%v,%T\n", y, y) // "65"

}
